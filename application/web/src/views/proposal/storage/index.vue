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
      prop="key"
      label="供给ID"
      width="120">
    </el-table-column>
    <el-table-column
      prop="user_id"
      label="仓储用户ID"
      width="120">
      <!-- <template slot-scope="scope">{{ scope.row.date }}</template> -->
    </el-table-column>
    <el-table-column
      prop="user_name"
      label="仓储用户名称"
      width="120">
    </el-table-column>
    <el-table-column
      prop="location_code"
      label="仓储地址编码"
      width="120">
    </el-table-column>
    <el-table-column
      prop="location_name"
      label="仓储地址名称"
      width="120">
    </el-table-column>
    <el-table-column
      prop="good_code"
      label="物资编码"
      width="120">
    </el-table-column>
    <el-table-column
      prop="good_name"
      label="物资名称"
      width="120">
    </el-table-column>
    <!-- <el-table-column
      prop="good_type"
      label="物资品类"
      width="120">
    </el-table-column> -->
    <el-table-column
      prop="count"
      label="仓储数量(箱)"
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
    <el-button @click="toTransport()">获取运力信息</el-button>
  </div>
  <!-- <el-pagination
  background
  layout="prev, pager, next"
  :total="1000">
  </el-pagination> -->
  </div>
</template>

<script>
  import { getStorageList } from '@/api/storage'

  export default {
    name: 'GenProposalStorage',
    data() {
      return {
        tableData: [],
        multipleSelection: []
      }
    },
    created() {
        getStorageList({
          pageSize: 10,
          bookmark: ""
        }).then( response => {
          if (response!==null) {
            // console.log(response)
            this.tableData = response.list
            // console.log(this.tableData)
          }
        }
        ).catch(err =>{console.log(error)})
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
      },
      toTransport() {
        this.$router.replace('/genProposalTransport')
      }
    }
  }
</script>