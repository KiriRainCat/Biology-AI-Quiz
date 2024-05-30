<template>
  <div>
    <!-- 完成 quiz 时，询问用户名的对话框 -->
    <var-dialog :show="quiz.showRecordNameDialog" :cancel-button="false" :confirm-button="false" width="360">
      <template #title>Saving to Leaderboard</template>
      <template #default>
        <p>Please enter your name/net-name/nickname to save your score to the leaderboard:</p>
        <var-input v-model="form.name" autofocus placeholder="Name"></var-input>
        <var-input v-model="form.passphrase" type="password" placeholder="Passphrase"></var-input>

        <var-button type="success" @click="() => quiz.uploadRecord(form.name, form.passphrase)" class="mt-5 w-full h-10">Save</var-button>
      </template>
    </var-dialog>

    <!-- 加载动画 -->
    <div v-if="!initialized" class="flex items-center justify-center h-screen">
      <var-loading type="cube">
        <template #description>Generating Quiz...</template>
      </var-loading>
    </div>

    <!-- 主视图 -->
    <div v-else class="flex h-screen items-center justify-center bg-gray-100">
      <div class="bg-white flex-none container relative overflow-hidden shadow-lg rounded-lg px-12 py-6">
        <div class="relative z-20 mt-8">
          <!-- 倒计时 -->
          <div class="bg-white shadow-lg p-1 rounded-full w-full h-5">
            <div class="bg-blue-700 rounded-full h-full" :style="`width: ${(quiz.countdown / COUNTDOWN_DURATION) * 100}%`"></div>
          </div>

          <!-- 问题 -->
          <div class="rounded-lg bg-gray-100 p-2 mt-5 neo-morph-1 text-center font-bold">
            <div class="bg-white p-5 text-gray-800">{{ quiz.currentQuestion.question }}</div>
          </div>

          <!-- 选项 -->
          <div class="mt-8">
            <div class="option" v-for="(option, i) in quiz.currentQuestion.choices" @click="quiz.nextQuestion(i)">
              <div class="rounded-lg p-3 w-full hover:bg-gray-50 bg-white">{{ option }}</div>
            </div>
          </div>

          <!-- 进度条 -->
          <div class="mt-8 text-center">
            <div class="h-1 w-12 bg-gray-800 rounded-full mx-auto"></div>
            <div class="font-bold text-sm text-gray-800">{{ quiz.currentQuestionIndex + 1 }} / {{ quiz!.questions.length }}</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { QuizController, COUNTDOWN_DURATION } from "@/pkg/services/quiz_logic";
import { Dialog } from "@varlet/ui";
import { onMounted, reactive, ref } from "vue";

let initialized = ref(false);
let quiz = ref(new QuizController());

onMounted(async () => {
  const isRetainedFromLocalStorage = await quiz.value.init();
  initialized.value = true;

  Dialog({
    title: `Quiz ${isRetainedFromLocalStorage ? "Retained" : "Generated"}`,
    message: "Ready to start? Let's go!!!",
    cancelButton: false,
    confirmButtonText: "Start!",
    onConfirm: () => quiz.value.startCountdown(),
    closeOnClickOverlay: false,
    closeOnKeyEscape: false,
  });
});

let form = reactive({ name: localStorage.getItem("name") || "", passphrase: localStorage.getItem("passphrase") || "" });
</script>
