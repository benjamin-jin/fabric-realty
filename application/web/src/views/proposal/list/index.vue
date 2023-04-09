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
      prop="admin_user_id"
      label="管理方用户ID"
      width="120">
    </el-table-column>
    <el-table-column
      prop="admin_user_name"
      label="管理方用户名"
      width="120">
      <!-- <template slot-scope="scope">{{ scope.row.date }}</template> -->
    </el-table-column>
    <el-table-column
      prop="proposal_duration"
      label="方案时长"
      width="120">
    </el-table-column>
    <el-table-column
      prop="proposal_cost"
      label="方案费用"
      width="120">
    </el-table-column>


    <!-- <el-table-column
      prop="good_name"
      label="物资名称"
      width="120">
    </el-table-column> -->
    <el-table-column
      prop="create_time"
      label="发布时间"
      width="120">
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
        <el-button
          type="text"
          @click="confirmProposal(scope.row)"
          size="small">
          确认
        </el-button>
      </template>
    </el-table-column>
  </el-table>
  <!-- <div style="margin-top: 20px">
    <el-button @click="toggleSelection()">取消选择</el-button>
    <el-button @click="toProposal()">发布方案</el-button>
  </div> -->
  <!-- <el-pagination
  background
  layout="prev, pager, next"
  :total="100">
  </el-pagination> -->
  </div>
</template>

<script>
  import { GetProposalList, DeleteProposal } from '@/api/proposal'

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
            GetProposalList({
              'user_id': "100001000001"
            }
             ).then( response => {
              if (response!==null) {
              // console.log(response)
                this.tableData = response
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
      confirmProposal(row) {
        DeleteProposal({
          'proposal_id': row.key
        }).then(
          res => {
            GetProposalList({
              'user_id': "100001000001"
            }
             ).then( response => {
              if (response!==null) {
              // console.log(response)
                this.tableData = response
              // console.log(this.tableData)
              }
            }
        ).catch(err =>{console.log(error)})
          }
        ).catch(err => {console.log(err)})
      }
      // toProposal() {
      //   PostProposal({
      //     proposal_id: this.proposalID,
      //     sub_proposal: JSON.stringify(this.tableData)
      //   }).then(res => {
      //     this.$router.replace('/genProposalList')
      //   }
      //   ).catch(err => consol.log(err))
      // },
      // getRowKeys(row) {
      //   return row.id
      // }
    }
  }
</script>