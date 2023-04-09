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
      label="运力ID"
      width="120">
    </el-table-column>
    <el-table-column
      prop="org_id"
      label="运力机构ID"
      width="120">
      <!-- <template slot-scope="scope">{{ scope.row.date }}</template> -->
    </el-table-column>
    <el-table-column
      prop="org_name"
      label="运力机构名称"
      width="120">
    </el-table-column>
    <el-table-column
      prop="start_location_code"
      label="供给地址编码"
      width="120">
    </el-table-column>
    <el-table-column
      prop="start_location_name"
      label="供给地址名称"
      width="120">
    </el-table-column>
    <el-table-column
      prop="end_location_code"
      label="需求地址编码"
      width="120">
    </el-table-column>
    <el-table-column
      prop="end_location_name"
      label="需求地址名称"
      width="120">
    </el-table-column>
    <el-table-column
      prop="transport_type"
      label="运力类型"
      width="120">
    </el-table-column>
    <!-- <el-table-column
      prop="good_name"
      label="物资名称"
      width="120">
    </el-table-column> -->
    <!-- <el-table-column
      prop="good_type"
      label="物资品类"
      width="120">
    </el-table-column> -->
    <el-table-column
      prop="transport_count"
      label="运力数量(辆)"
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
    <el-button @click="toPost()">生成方案</el-button>
  </div>
  <!-- <el-pagination
  background
  layout="prev, pager, next"
  :total="100">
  </el-pagination> -->
  </div>
</template>

<script>
  import { getTransportList } from '@/api/transport'

  export default {
    name: 'GenProposalTransport',
    data() {
      return {
        tableData: [],
        multipleSelection: [],
        startLocationCodes: [
          '330101',
          '330102',
          '330103',
          '330104',
        ],
        endLocationCodes: [
          '330105',
          '330106',
          '330107'
        ],
        currentPage: 1,
        pageSize: 10
      }
    },
    created() {
        for (const slc of this.startLocationCodes) {
          for (const elc of this.endLocationCodes) {
            getTransportList({
              start_location_code: slc,
              end_location_code: elc
            }).then( response => {
              if (response!==null) {
              // console.log(response)
              this.tableData = this.tableData.concat(response)
              // console.log(this.tableData)
              }
            }
        ).catch(err =>{console.log(error)})
          }
        }

    },

    methods: {
      // currentChange(currentPage) {
      //   this.currentPage = currentPage
      //   console.log(this.currentPage)
      // },
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
      toPost() {
        this.$router.replace('/genProposalPost')
      },
      // getRowKeys(row) {
      //   return row.id
      // }
    }
  }
</script>