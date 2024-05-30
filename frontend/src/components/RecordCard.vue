<script setup lang="ts">
import type { Record } from "@/pkg/services/api_types";
import type { PropType } from "vue";

defineProps({ record: Object as PropType<Record>, onlyData: Boolean });
</script>

<template>
  <var-paper elevation="1" radius="6" class="flex gap-6 p-3 bg-cyan-50">
    <div v-if="!onlyData">{{ record!.name }}</div>
    <div>{{ (record!.accuracy * 100).toFixed(0) }}%</div>
    <div>{{ record!.time_taken.toFixed(1) }}s</div>
    <div v-if="!onlyData">{{ new Date(record!.updated_at!).toLocaleString() }}</div>
  </var-paper>
</template>
