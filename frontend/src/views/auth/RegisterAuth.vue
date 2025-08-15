<script setup lang="ts">
import router from '@/router';
import axios from 'axios';
import { reactive } from 'vue';
    const reg = reactive({
        Username: "",
        Email: "",
        Password: "",
    })

    const submitReg = async() => {
        await axios.post('/api/register', reg).then(response => {
                const resp = response.data
                console.log(resp['message'], response.status);
                if (response.status === 201 ) {
                    router.push('/login')
                }
            }).catch(err =>{
                console.log(err);
                
            })
    }
</script>

<template>
    Ini Halaman Register

    <nav>
        <RouterLink to="/">Home</RouterLink>
    </nav>

    <form  @submit.prevent="submitReg">
        <label for="username">Username</label>
        <input type="text" id="username" v-model="reg.Username">
        <label for="email">Email</label>
        <input type="email" id="email" v-model="reg.Email">
        <label for="password">Password</label>
        <input type="text" id="password" v-model="reg.Password">
        <button type="submit">Register</button>
    </form>
</template>