<script setup lang="ts">
import router from '@/router';
import axios from 'axios';
import { reactive } from 'vue';

    const reg = reactive({
        Email:"",
        Password:""
    })

    const submitLogin = async() => {
        await axios.post("/api/login", reg).then(response =>{
            console.log(response.data);

            // set localstorage
            localStorage.setItem("token", "Bearer " + response.data.token)
            router.push("/dashboard")

        }).catch(err => {
            console.log(err.response.data.message);
            router.back
        })
    }

</script>

<template>
    ini halaman login

    <form @submit.prevent="submitLogin">
        <label for="email">Email</label>
        <input type="email" id="email" v-model="reg.Email">
        <label for="password">Password</label>
        <input type="text" id="password" v-model="reg.Password">

        <button type="submit">login</button>

    </form>
</template>