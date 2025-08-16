<script setup lang="ts">
import router from '@/router';
import axios from 'axios';

axios.get('/api/admin/role', {
    headers: {
        Authorization: localStorage.getItem("token")
    }
}).then(function (response) {
    console.log(response);

}).catch(function (error) {
    console.log(error);
    router.push("/login")
});

const tes = async() =>{
    axios.get("/api/admin/logout",{
        headers: {
            Authorization: localStorage.getItem("token")
        }
    }).then(res => {
        console.log(res.data.message);
        localStorage.removeItem("token");

        router.push("/")
    }).catch(err =>{
        console.log(err);
    })
} 
</script>

<template>
    <h1>ini halaman admin</h1>

    <form @submit.prevent="tes">
        <button type="submit">Logout</button>
    </form>
    
</template>