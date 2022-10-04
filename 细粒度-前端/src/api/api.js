let host = "/api";
import axios from "axios";
import ElementUI from 'element-ui'

export default {
  generateRsaKeys: host + "/common/rsa",
  DABEUser: host + "/dabe/user",
  platUser: host + "/user/",
  DABEUserAttr: host + "/dabe/user/attr",
  platUserAttr: host + "/user/attr",
  orgApply: host + "/org/apply",
  org: host + "/org/",
  applyOthersAttr: host + "/user/attr/apply",
  approveAttrApply: host + "/user/attr/approval",
  syncAttr: host + "/user/attr/sync",
  encryptAndUpload: host + "/content/",
  getContents: host + "/content/list",
  decryptContent: host + "/content/decryption",

  applyCreateOrg: host + "/org/apply/creation",
  approveJoinOrg: host + "/org/apply/creation/approval",
  sharePKForOrg: host + "/org/part-pk",
  completePK: host + "/org/complete-pk",

  applyCreateOrgAttr: host + "/org/apply/attribute",
  approveOrgAttr: host + "/org/apply/attribute/approval",
};

export function post(url, params) {
  return new Promise((resolve, reject) => {
    axios.post(url, params).then((res) => {
      if (res.data.code === 200) {
        resolve(res);
      } else {
        throw res
      }
    }).catch((err) => {
      ElementUI.Message(err.data.message)
      reject(err);
    });
  });
}

export function get(url, params) {
  return new Promise((resolve, reject) => {
    axios.get(url, params).then((res) => {
      if (res.data.code === 200) {
        resolve(res);
      } else {
        throw res
      }
    }).catch((err) => {
      ElementUI.Message(err.data.message)
      reject(err);
    });
  });
}