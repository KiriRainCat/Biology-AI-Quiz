<script setup lang="ts">
import type { Question } from "@/pkg/services/api_types";
import { ref } from "vue";

const questions = ref<Question[]>(JSON.parse(localStorage.getItem("prevQuestions") ?? "[]"));
const selectedOptions = ref<number[]>(JSON.parse(localStorage.getItem("selectedOptions") ?? "[]"));

const currentIdx = ref(0);
</script>

<template>
  <div class="flex h-screen items-center justify-center bg-gray-100">
    <div class="bg-white flex-none container relative overflow-hidden shadow-lg rounded-lg px-12 py-6">
      <div class="relative z-20 mt-8">
        <!-- 问题 -->
        <div class="rounded-lg bg-gray-100 p-2 mt-5 neo-morph-1 text-center font-bold">
          <div class="bg-white p-5 text-gray-800">{{ questions[currentIdx].question }}</div>
        </div>

        <!-- 选项 -->
        <div class="mt-8">
          <div class="option" v-for="(option, i) in questions[currentIdx].choices">
            <div
              class="rounded-lg p-3 w-ful"
              :class="`${i == questions[currentIdx].answer ? 'bg-green-300' : i == selectedOptions[i] && questions[currentIdx].answer != selectedOptions[i] ? 'bg-red-300' : 'bg-white'}`"
            >
              {{ option }}
            </div>
          </div>
        </div>

        <!-- 进度条 -->
        <div class="mt-8 flex justify-center gap-3">
          <var-button text size="mini" class="text-base" :disabled="currentIdx <= 0" @click="() => currentIdx--">←</var-button>
          <div class="text-center">
            <div class="h-1 w-12 bg-gray-800 rounded-full mx-auto"></div>
            <div class="font-bold text-sm text-gray-800">{{ currentIdx + 1 }} / {{ questions.length }}</div>
          </div>
          <var-button text size="mini" class="text-base" :disabled="currentIdx >= questions.length - 1" @click="currentIdx++">→</var-button>
        </div>
      </div>
    </div>
  </div>
</template>
