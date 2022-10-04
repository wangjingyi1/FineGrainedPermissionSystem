import api from "./api";
import { post, get } from "./api";
import qs from "qs"

export function encryptAndUpload(fileName, tags, plainContent, policy) {
    return post(api.encryptAndUpload, {
        "fileName": fileName, "tags": tags, "plainContent": plainContent, "policy": policy,
    })
}

export function getContents(fromUserName, tag, pageSize, bookmark) {
    return get(api.getContents, {
        params: {
            "fromUserName": fromUserName, "tag": tag, "pageSize": pageSize, "bookmark": bookmark,
        }
    })
}

export function decryptContent(fileName, cipher) {
    return post(api.decryptContent, {
        "fileName": fileName, "cipher": cipher
    })
}