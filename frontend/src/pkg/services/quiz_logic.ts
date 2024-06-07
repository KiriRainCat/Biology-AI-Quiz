import { Snackbar } from "@varlet/ui";
import router from "../router";
import { addRecord, generateQuiz } from "./api";
import type { Question } from "./api_types";

export const COUNTDOWN_DURATION = (36 * 1000) / 60;

export class QuizController {
  questions: Question[] = [];
  currentQuestionIndex: number = 0;
  public get currentQuestion(): Question {
    return this.questions[this.currentQuestionIndex];
  }

  selectedOptions: number[] = [];

  /** 单位为秒 */
  timeTaken: number = 0;
  numOfCorrect: number = 0;

  countdown: number = COUNTDOWN_DURATION;
  timer?: number;

  showRecordNameDialog: boolean = false;

  /**
   * 初始化，获取题目
   * @returns 是否是本地 localStorage 缓存读取的题目
   */
  public async init(): Promise<boolean> {
    localStorage.removeItem("prevQuestions");

    // 尝试从本地缓存读取未做过的题目
    const questionsStr = localStorage.getItem("questions");
    if (questionsStr && questionsStr.length > 60) {
      this.questions = JSON.parse(questionsStr);
      return true;
    }

    // 否则从 API 获取题目
    this.questions = await generateQuiz();
    if (this.questions.length == 0) {
      router.push("/");
      Snackbar.error("Failed to generate, please try again.");
      return false;
    }

    // 缓存题目
    localStorage.setItem("questions", JSON.stringify(this.questions));
    if (this.questions.length == 0) router.push("/");
    return false;
  }

  /**
   * 开始计时器
   *
   * 每次调用都会重置计时器
   */
  public startCountdown() {
    clearInterval(this.timer);
    this.timeTaken += ((COUNTDOWN_DURATION - this.countdown) * 60) / 1000;
    this.countdown = COUNTDOWN_DURATION;

    this.timer = setInterval(() => {
      if (this.countdown <= 0) {
        clearInterval(this.timer);
        this.nextQuestion(-1);
        return;
      }

      this.countdown--;
    }, 60);
  }

  /**
   * 提交答案，转至下一题
   * @param selectedOptionIdx 选项索引
   */
  public nextQuestion(selectedOptionIdx: number) {
    // 记录答案
    this.selectedOptions[this.currentQuestionIndex] = selectedOptionIdx;
    localStorage.setItem("selectedOptions", JSON.stringify(this.selectedOptions));

    // 判断答案是否正确
    if (selectedOptionIdx == this.questions[this.currentQuestionIndex].answer) this.numOfCorrect++;

    // 判断是否还有题目，无题则结束
    if (this.currentQuestionIndex >= this.questions.length - 1) {
      clearInterval(this.timer);

      // 清除缓存的题目
      let questionsStr = localStorage.getItem("questions");
      localStorage.removeItem("questions");
      localStorage.setItem("prevQuestions", questionsStr ?? "");

      // 询问在 leaderboard 中显示的名称
      this.showRecordNameDialog = true;
      return;
    }

    // 转至下一题
    this.currentQuestionIndex++;
    this.startCountdown();
  }

  /**
   * 上传 quiz 结果
   */
  public async uploadRecord(name: string, passphrase: string) {
    let isSuccess = await addRecord({
      name: name,
      accuracy: this.numOfCorrect / this.questions.length,
      time_taken: this.timeTaken,
      passphrase: passphrase,
    });

    if (isSuccess) {
      // 将 name 和 passphrase 保存到本地缓存
      localStorage.setItem("name", name);
      localStorage.setItem("passphrase", passphrase);

      Snackbar.success("Record uploaded successfully!");
      this.goToResult();
    }
  }

  public goToResult() {
    this.showRecordNameDialog = false;
    router.push("/result");
  }
}
