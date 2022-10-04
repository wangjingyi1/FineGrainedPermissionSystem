import api from "./api";
import { post, get } from "./api";
import qs from "qs"


export function getOrgApply(orgName, type) {
    return get(api.orgApply, { params: { "orgName": orgName, "type": type } })
}

export function getOrgAttrApply(orgName, type, attrName) {
    return get(api.orgApply, { params: { "orgName": orgName, "type": type, "attrName": attrName } })
}

export function getOrgInfo(orgName) {
    return get(api.org, { params: { "orgName": orgName } })
}

export function applyCreateOrg(fileName, t, n, users, orgName,) {
    return post(api.applyCreateOrg, {
        "fileName": fileName, "t": t, "n": n, "users": users, "orgName": orgName,
    })
}

export function approveJoinOrg(fileName, orgName, attrName) {
    return post(api.approveJoinOrg, {
        "fileName": fileName, "attrName": attrName, "orgName": orgName,
    })
}

export function sharePkForOrg(type, orgName, fileName, attrName) {
    return post(api.sharePKForOrg, qs.stringify({
        "fileName": fileName, "type": type, "orgName": orgName, "attrName": attrName,
    }))
}

export function completePK(type, orgName, fileName, attrName) {
    return post(api.completePK, qs.stringify({
        "fileName": fileName, "type": type, "orgName": orgName, "attrName": attrName,
    }))
}


export function applyCreateOrgAttr(fileName, attrName, orgName,) {
    return post(api.applyCreateOrgAttr, {
        "fileName": fileName, "attrName": attrName, "orgName": orgName,
    })
}

export function approveOrgAttr(fileName, orgName, attrName) {
    return post(api.approveOrgAttr, {
        "fileName": fileName, "attrName": attrName, "orgName": orgName,
    })
}