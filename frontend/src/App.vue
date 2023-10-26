<template>
  <div class="common-layout">
    <el-container>
      <el-header class="header" >Cave_CTF</el-header>
      <el-container>
        <el-aside class="aside" width="200px">
          <el-row class="tac">
            <el-col :span="24">
              <h5 style="text-align: left;padding-left: 25%" class="mb-2">导航栏</h5>
              <el-menu
                  default-active="1"
                  class="el-menu-vertical-demo"
                  @open="handleOpen"
                  @close="handleClose"
              >
                <el-menu-item index="1">
                  <el-icon><icon-menu /></el-icon>
                  <span>题目列表</span>
                </el-menu-item>
                <el-menu-item index="2">
                  <el-icon><document /></el-icon>
                  <span>分数排行</span>
                </el-menu-item>
                <el-menu-item index="3">
                  <el-icon><setting /></el-icon>
                  <span>成员列表</span>
                </el-menu-item>
              </el-menu>
            </el-col>
          </el-row>
        </el-aside>
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
