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
    <div v-if="loginSwitch">
      <div class="form-container">
        <!-- 登录表单 -->
        <div class="form">
          <h3>Cave-CTF</h3>
          <el-input v-model="login" class="input" placeholder="请输入用户名" />
          <el-input v-model="password" class="input" type="password" placeholder="请输入密码" />
          <el-button class="btn" @click="login">登录</el-button>
        </div>
      </div>
    </div>
    <div v-else>
      <div class="form-container">
        <!-- 注册表单 -->
        <div class="form">
          <h3>Cave-CTF</h3>
          <el-input v-model="name" class="input" placeholder="请输入用户名" />
          <el-input v-model="email" class="input" placeholder="请输入邮箱" />
          <el-input v-model="password" class="input" type="password" placeholder="请输入密码" />
          <el-button class="btn" @click="register">注册</el-button>
        </div>
      </div>
    </div>
    <div>
      <el-button @click="toggleForm">{{ loginSwitch ? '切换到注册' : '切换到登录' }}</el-button>
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
export default  {
  data(){
    return{
      login: "",
      name: "",
      pass:"",
      activeMenuItem:1,
      loginStatus:false,
      loginSwitch:true,
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
  }
}
</script>