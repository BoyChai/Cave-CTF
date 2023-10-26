<template>
  <div class="common-layout">
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
</style>

<script>
export default  {
  data(){
    return{
      activeMenuItem:1,
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