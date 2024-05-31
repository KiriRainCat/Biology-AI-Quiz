<script setup lang="ts">
import type { Record } from "@/pkg/services/api_types";
import type { PropType } from "vue";

defineProps({ record: Object as PropType<Record>, onlyData: Boolean });
</script>

<template>
  <tr>
    <td v-if="!onlyData">{{ record!.name }}</td>
    <td>{{ (record!.accuracy * 100).toFixed(0) }}%</td>
    <td>{{ record!.time_taken.toFixed(1) }}s</td>
    <td v-if="!onlyData">{{ new Date(record!.updated_at!).toLocaleString() }}</td>
  </tr>
</template>
