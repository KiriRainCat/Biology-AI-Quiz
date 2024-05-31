<script setup lang="ts">
import RecordCard from "@/components/RecordCard.vue";
import { getRecords } from "@/pkg/services/api";
import type { Record } from "@/pkg/services/api_types";
import { onBeforeUnmount, onMounted, ref } from "vue";

let autoFetchInterval: number;
const records = ref<Record[]>([]);

onMounted(async () => {
  records.value = await getRecords();
  autoFetchInterval = setInterval(async () => {
    records.value = await getRecords();
  }, 10000);
});

onBeforeUnmount(() => clearInterval(autoFetchInterval));
</script>

<template>
  <div class="flex flex-col items-center justify-center px-6">
    <div class="text-2xl font-bold mt-6">Leaderboard</div>
    <div class="mb-6 text-xs opacity-50">Data will auto refetch every 10 seconds</div>

    <!-- 加载动画 -->
    <var-loading v-if="records.length == 0" />

    <!-- 排行榜主体 -->
    <var-table v-else elevation="0" scroller-height="76vh" class="max-w-[46rem]">
      <thead class="font-bold">
        <tr>
          <td>Name</td>
          <td>Acc.</td>
          <td>Time</td>
          <td>Recorded On</td>
        </tr>
      </thead>
      <RecordCard v-for="record in records" :key="record.name" :record="record" class="mb-1.5" />
    </var-table>

    <!-- 返回按钮 -->
    <var-button type="primary" @click="() => $router.back()" class="absolute bottom-4">Back</var-button>
  </div>
</template>
