<template>
  <div>
    <!--    axiosOutput获取到的后端数据用来展示-->
    axios output:{{axiosOutput}}
  </div>
</template>

<script setup>
// 获取后端的数据
import axios from "axios";
const axiosOutput=ref()
axios.get('http://localhost:8000/test',{
  // params:{
  //   type:'pop',page:1
  // },
}).then(res=>{
  console.log(res);
  axiosOutput.value=res['data'] // 把获得的后端数据res，赋值给axiosOutput，用来在前端展示
}).catch(err=>{
  console.log(err);
})
import {ref} from "vue";

</script>
