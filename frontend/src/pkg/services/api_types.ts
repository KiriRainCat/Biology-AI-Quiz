/** 问题数据模型 */
interface Question {
  question: string;
  choices: string[];
  answer: number;
}

/** 用户 quiz 结果数据模型 */
interface Record {
  name: string;
  passphrase?: string;
  accuracy: number;
  time_taken: number;
  updated_at?: string;
}

export type { Question, Record };
