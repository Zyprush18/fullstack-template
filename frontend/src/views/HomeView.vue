<script setup lang="ts">
import axios from 'axios';
import TheWelcome from '../components/TheWelcome.vue'

axios.get('/api/role').then(function (response) {
  console.log(response);
  
}).catch(function (error) {
  console.log(error);
  
})


</script>

<template>
  <main>
    <TheWelcome />
  </main>
</template>
