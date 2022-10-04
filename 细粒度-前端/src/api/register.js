import api from "./api";
import { post, get } from "./api";
import qs from "qs"


export function generateRsaKeys(fileName) {
    return post(api.generateRsaKeys, qs.stringify({ "fileName": fileName }));
}

export function createDABEUser(fileName, userName) {
    return post(api.DABEUser, qs.stringify({ "fileName": fileName, "userName": userName }))
}

export function getDABEUser(fileName) {
    return get(api.DABEUser, { params: { "fileName": fileName } })
}

export function platUser(fileName) {
    return post(api.platUser, { "fileName": fileName })
}