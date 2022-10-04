<template>
  <el-container style="height: 900px; border: 1px solid #eee">
    <el-container>
      <el-header>用户注册</el-header>

      <el-container style="height: 0px; border: 1px">
        <el-aside>文件名</el-aside>
        <el-main>
          <el-input placeholder="存放在本机的文件名，建议与GID相同"
                    v-model="fileName"
                    clearable>
          </el-input>
        </el-main>
      </el-container>

      <el-container style="height: 0px; border: 1px">
        <el-aside>用户名</el-aside>
        <el-main>
          <el-input placeholder="即GID"
                    v-model="userName"
                    clearable>
          </el-input>
        </el-main>
      </el-container>

      <el-container>
        <el-aside>注册步骤</el-aside>
        <el-main>
          <el-container style="height: 300px"
                        v-show="active === 0">
            <el-main>
              <el-row>
                <el-col :span="4">
                  <div>私钥:</div>
                </el-col>
                <el-col :span="20">
                  <textarea v-model="priKey"
                            disabled=true
                            id="key"></textarea>
                </el-col>
              </el-row>

              <el-row>
                <el-col :span="4">
                  <div>公钥:</div>
                </el-col>
                <el-col :span="20">
                  <textarea v-model="pubKey"
                            disabled=true
                            id="key"></textarea>
                </el-col>
              </el-row>

              <el-row>
                <el-button style="margin-top: 12px;"
                           @click="generateKeys">生成公私钥对</el-button>
              </el-row>
            </el-main>
          </el-container>

          <el-container style="height: 300px"
                        v-show="active === 1">
            <el-main>
              <el-row>
                <el-col :span="4">
                  <div>属性私钥:</div>
                </el-col>
                <el-col :span="20">
                  <textarea v-model="abePriKey"
                            disabled=true
                            id="key"></textarea>
                </el-col>
              </el-row>

              <el-row>
                <el-col :span="4">
                  <div>属性公钥:</div>
                </el-col>
                <el-col :span="20">
                  <textarea v-model="abePubKey"
                            disabled=true
                            id="key"></textarea>
                </el-col>
              </el-row>

              <el-row>
                <el-button style="margin-top: 12px;"
                           @click="generateKeys2">DABE生成属性公私钥对</el-button>
              </el-row>
            </el-main>
          </el-container>

          <el-container style="height: 300px"
                        v-show="active === 2">
            <el-main>
              <el-row>
                <el-col :span="4">
                  <div>私钥:</div>
                </el-col>
                <el-col :span="20">
                  <textarea v-model="priKey"
                            disabled=true
                            id="key"></textarea>
                </el-col>
              </el-row>

              <el-row>
                <el-col :span="4">
                  <div>公钥:</div>
                </el-col>
                <el-col :span="20">
                  <textarea v-model="pubKey"
                            disabled=true
                            id="key"></textarea>
                </el-col>
              </el-row>
              <el-row>
                <el-col :span="4">
                  <div>属性私钥:</div>
                </el-col>
                <el-col :span="20">
                  <textarea v-model="abePriKey"
                            disabled=true
                            id="key"></textarea>
                </el-col>
              </el-row>

              <el-row>
                <el-col :span="4">
                  <div>属性公钥:</div>
                </el-col>
                <el-col :span="20">
                  <textarea v-model="abePubKey"
                            disabled=true
                            id="key"></textarea>
                </el-col>
              </el-row>
            </el-main>
          </el-container>

          <el-steps :active="active"
                    finish-status="success"
                    align-center>
            <el-step title="step 1"></el-step>
            <el-step title="step 2"></el-step>
            <el-step title="step 3"></el-step>
          </el-steps>

          <el-button style="margin-top: 12px;"
                     @click="next">{{description}}</el-button>

        </el-main>
      </el-container>

    </el-container>
  </el-container>
</template> 


<script>
  import { createDABEUser, generateRsaKeys, getDABEUser, platUser } from "../../api/register";

  export default {
    data() {
      return {
        fileName: '',
        userName: '',
        active: 0,
        description: '下一步',
        priKey: '',
        pubKey: '',
        abePriKey: '',
        abePubKey: '',
      }
    },
    methods: {
      next() {
        this.active++
        if (this.active === 2) {
          this.description = '提交'
        }
        if (this.active === 3) {
          this.register()
        }
      },
      generateKeys() {
        if (this.fileName === '') {
          this.$message("fileName 为空")
          return
        }
        generateRsaKeys(this.fileName).then(res => {
          this.priKey = res.data.data.priKey
          this.pubKey = res.data.data.pubKey
        })
      },
      generateKeys2() {
        if (this.fileName === '') {
          this.$message("fileName 为空")
          return
        }
        if (this.userName === '') {
          this.$message("userName 为空")
          return
        }

        createDABEUser(this.fileName, this.userName).then(res => {
          this.abePriKey = res.data.data.Alpha
          this.abePubKey = res.data.data.EGGAlpha
        }).catch(err => {
          console.log(err)
          getDABEUser(this.fileName).then(res => {
            this.abePriKey = res.data.data.Alpha
            this.abePubKey = res.data.data.EGGAlpha
          })
        })
        
      },
      register() {
        platUser(this.fileName).then(res => {
          this.$message('注册成功');
          this.$router.push({
            name: "RegisterLogin"
          });
        })
        
      },
    }
  };
</script>

<style>
.el-header,
.el-footer {
  background-color: #545c64;
  color: #ffffff;
  text-align: center;
  line-height: 60px;
}

.el-aside {
  background-color: #545c64;
  color: #ffffff;
  /* text-align: center; */
  line-height: 80px;
}

.el-main {
  background-color: #ffffff;
  color: #333;
  text-align: center;
}

body > .el-container {
  margin-bottom: 40px;
}

.el-container:nth-child(5) .el-aside,
.el-container:nth-child(6) .el-aside {
  line-height: 260px;
}

.el-container:nth-child(7) .el-aside {
  line-height: 320px;
}
.el-row {
  margin-bottom: 20px;
  height: 30px;
  &:last-child {
    margin-bottom: 0;
  }
}
.el-col {
  border-radius: 4px;
}
#key {
  width: 100%;
  resize: none;
}
</style>