<template>
  <div style="margin-top:20px;margin-left: 100px;margin-right: 100px;">
   <el-table :data="data" style="width: 100%">
     <el-table-column prop="ID" label="ID" width="180" />
     <el-table-column prop="Name" label="Name" width="180" />
     <el-table-column prop="Score" label="Score" />
    </el-table>
  </div>
</template>

<script>
export default {
  name: "users",
  data() {
    return {
      data:[{"ID":1,"Name":"test1","Score":"0"},{"ID":2,"Name":"tset2","Score":"1"},{"ID":3,"Name":"test3","Score":"1"}],
    }
  }
}
</script>

<style scoped>

</style>