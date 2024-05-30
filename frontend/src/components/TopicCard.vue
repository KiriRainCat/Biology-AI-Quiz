<script setup lang="ts">
defineProps({
  topic: String,
});
</script>

<template>
  <var-paper elevation="1" class="w-min px-2 py-1 rounded-md bg-emerald-100">{{ topic }}</var-paper>
</template>
