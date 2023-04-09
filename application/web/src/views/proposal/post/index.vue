<template>
  <div class="app-container">
  <el-table
    ref="multipleTable"
    :data="tableData"
    tooltip-effect="dark"
    style="width: 100%">

    <!-- <el-table-column
      type="selection"
      width="55">
    </el-table-column> -->
    <el-table-column
      prop="storage_user_id"
      label="供应方用户ID"
      width="120">
    </el-table-column>
    <el-table-column
      prop="storage_user_name"
      label="供应方用户名"
      width="120">
      <!-- <template slot-scope="scope">{{ scope.row.date }}</template> -->
    </el-table-column>
    <el-table-column
      prop="transport_org_id"
      label="运力机构ID"
      width="120">
    </el-table-column>
    <el-table-column
      prop="transport_org_name"
      label="运力机构名称"
      width="120">
    </el-table-column>
    <el-table-column
      prop="start_location_code"
      label="仓储地址编码"
      width="120">
    </el-table-column>
    <el-table-column
      prop="start_location_name"
      label="仓储地址名称"
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
    <el-table-column
      prop="transport_count"
      label="运力数量"
      width="120">
    </el-table-column>
    <el-table-column
      prop="sub_proposal_duration"
      label="子方案时长(小时*箱)"
      width="120">
    </el-table-column>
    <el-table-column
      prop="sub_proposal_cost"
      label="子方案总费用(元)"
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
      fixed="right"
      label="操作"
      width="120">
      <template slot-scope="scope">
        <el-button
          type="text"
          size="small">
          +
        </el-button>
        <el-button
          type="text"
          size="small">
          -
        </el-button>
      </template>
    </el-table-column>
  </el-table>
  <div style="margin-top: 20px">
    <el-button @click="toggleSelection()">取消选择</el-button>
    <el-button @click="toProposal()">发布方案</el-button>
  </div>
  <!-- <el-pagination
  background
  layout="prev, pager, next"
  :total="100">
  </el-pagination> -->
  </div>
</template>

<script>
  import { CreateProposal, GetProposal, PostProposal } from '@/api/proposal'

  export default {
    name: 'GenProposalTransport',
    data() {
      return {
        proposalID: "",
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
        pageSize: 10,
        cost: 0,
        duration: 0
      }
    },
    created() {
            CreateProposal(
             ).then( response => {
              if (response!==null) {
              // console.log(response)
                this.proposalID = response.proposal_id
                GetProposal({'proposal_id':response.proposal_id}).then(response => {
                  if (response!==null) {
                    console.log(response)
                    this.tableData = response.alg_sub_list
                    this.cost = response.proposal_cost
                    this.duration = response.proposal_duration
                  }
                }
                ).catch(err =>{console.log(error)})
              // console.log(this.tableData)
              }
            }
        ).catch(err =>{console.log(error)})

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
      toProposal() {
        PostProposal({
          proposal_id: this.proposalID,
          sub_proposal: JSON.stringify(this.tableData)
        }).then(res => {
          this.$router.replace('/genProposalList')
        }
        ).catch(err => consol.log(err))
      },
      // getRowKeys(row) {
      //   return row.id
      // }
    }
  }
</script>