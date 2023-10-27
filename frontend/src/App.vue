<template>
  <div v-if="loginStatus" class="common-layout">
    <el-container>
<!--      顶部Header-->
      <el-header class="header" >
        <br>
        Cave_CTF
      </el-header>
      <el-container>
<!--        左侧导航栏-->
        <el-aside class="aside" width="200px">
          <el-row class="tac">
            <el-col :span="24">
              <h5 style="text-align: left;padding-left: 25%" class="mb-2">导航栏</h5>
              <el-menu
                  :default-active="activeMenuItem"
                  class="el-menu-vertical-demo"
              >
                <el-menu-item @click="question" index="1">
                  <el-icon><icon-menu /></el-icon>
                  <span>题目列表</span>
                </el-menu-item>
                <el-menu-item @click="ranking" index="2">
                  <el-icon><document /></el-icon>
                  <span>分数排行</span>
                </el-menu-item>
                <el-menu-item @click="users"  index="3">
                  <el-icon><setting /></el-icon>
                  <span>成员列表</span>
                </el-menu-item>
              </el-menu>
            </el-col>
          </el-row>
        </el-aside>
<!--        主要内容-->
        <el-main class="main">
          <router-view/>
        </el-main>
      </el-container>
    </el-container>
  </div>

  <div v-if="!loginStatus">
    <div class="form-container">
      <!-- 注册表单 -->
      <div class="form">
        <h3>Cave-CTF</h3>
        <el-input v-model="name" class="input" placeholder="请输入真实姓名" />
        <el-input v-model="alias" class="input" placeholder="请输入别名" />
        <el-button class="btn" @click="login">登录/注册</el-button>
      </div>
    </div>
    </div>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}
.header {
  border: 1px rebeccapurple solid;
}
.aside{
  border: 1px royalblue solid;
}
.main{
  border: 1px red solid;
  height: 650px;
}
.form-container {
   display: flex;
   justify-content: center;
   align-items: center;
   height: 90vh;
 }

.form {
  text-align: center;
  width: 300px;
  border: 1px solid #ccc;
  padding: 20px;
  border-radius: 5px;
  background-color: #f0f0f0;
}

.input {
  width: 100%;
  margin-bottom: 10px;
}

.btn {
  width: 100%;
}
</style>

<script>
import axios from 'axios';

export default  {
  data(){
    return{
      name: "",
      alias:"",
      activeMenuItem:1,
      loginStatus:false,
    }
  },
  methods: {
    question() {
      this.$router.push('/question')
      this.activeMenuItem=1
    },
    ranking(){
      this.$router.push('/ranking')
      this.activeMenuItem=2
    },
    users(){
      this.$router.push('/users')
      this.activeMenuItem=3
    },
    login() {
        if (this.name === '' || this.alias === '') {
          this.$message({
            message: '用户名和别名不允许为空',
            type: 'warning'
          });
          return
        }
        axios.get("/check/user",{
          params:{
            name: this.name,
            alias: this.alias
          }
        }).then(
            res => {
              //请求成功执行
              this.loginStatus=true
              document.cookie = "name="+this.name+"; alias="+this.alias
            }
        ).catch(error => {
          //请求失败执行
          this.$message({
            message: '用户未注册,帮您创建用户',
            type: 'info'
          });
          axios.post("/create/user",{
            name:this.name,
            alias: this.alias
          }).then(r=>{
            this.loginStatus = true
            document.cookie = "name="+this.name+"; alias="+this.alias
          }).catch(e =>{
            this.$message({
              message:"登录/注册失败，请联系管理员",
              type:"error",
            })
          })
        })
    },
  },
  created() {
    if (document.cookie === '') {
      return
    }
    this.loginStatus = true
    console.log(document.cookie)
  }
}
</script>