<template>
  <div class="app-container">
  <el-table
    ref="multipleTable"
    :data="tableData"
    tooltip-effect="dark"
    style="width: 100%"
    @selection-change="handleSelectionChange">
    <el-table-column
      type="selection"
      width="55">
    </el-table-column>
    <el-table-column
      prop="demandID"
      label="需求ID"
      width="120">
    </el-table-column>
    <el-table-column
      prop="userID"
      label="用户ID"
      width="120">
      <!-- <template slot-scope="scope">{{ scope.row.date }}</template> -->
    </el-table-column>
    <el-table-column
      prop="userName"
      label="用户名称"
      width="120">
    </el-table-column>
    <el-table-column
      prop="locationCode"
      label="需求地址编码"
      width="120">
    </el-table-column>
    <el-table-column
      prop="locationName"
      label="需求地址名称"
      width="120">
    </el-table-column>
    <el-table-column
      prop="goodCode"
      label="物资编码"
      width="120">
    </el-table-column>
    <el-table-column
      prop="goodName"
      label="物资名称"
      width="120">
    </el-table-column>
    <el-table-column
      prop="goodType"
      label="物资品类"
      width="120">
    </el-table-column>
    <el-table-column
      prop="count"
      label="需求数量(箱)"
      show-overflow-tooltip>
    </el-table-column>
    <el-table-column
      fixed="right"
      label="操作"
      width="120">
      <template slot-scope="scope">
        <el-button
          type="text"
          size="small">
          查看详情
        </el-button>
      </template>
    </el-table-column>
  </el-table>
  <div style="margin-top: 20px">
    <el-button @click="toggleSelection()">取消选择</el-button>
    <el-button @click="">获取仓储信息</el-button>
  </div>
  <el-pagination
  background
  layout="prev, pager, next"
  :total="1000">
  </el-pagination>
  </div>
</template>

<script>
  export default {
    name: 'GenProposal',
    data() {
      return {
        tableData: [],
        multipleSelection: []
      }
    },

    methods: {
      toggleSelection(rows) {
        if (rows) {
          rows.forEach(row => {
            this.$refs.multipleTable.toggleRowSelection(row);
          });
        } else {
          this.$refs.multipleTable.clearSelection();
        }
      },
      handleSelectionChange(val) {
        this.multipleSelection = val;
      }
    }
  }
</script>