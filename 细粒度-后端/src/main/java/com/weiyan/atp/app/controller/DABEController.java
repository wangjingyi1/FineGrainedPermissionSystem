package com.weiyan.atp.app.controller;

import com.weiyan.atp.data.bean.DABEUser;
import com.weiyan.atp.data.bean.Result;
import com.weiyan.atp.service.DABEService;

import lombok.extern.slf4j.Slf4j;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author : 魏延thor
 * @since : 2020/6/2
 */
@RestController
@Slf4j
@RequestMapping("/dabe")
public class DABEController {
    private final DABEService dabeService;

    public DABEController(DABEService dabeService) {
        this.dabeService = dabeService;
    }

    @GetMapping("/user")
    public Result<DABEUser> getUser(String fileName) {
        return handleUser(dabeService.getUser(fileName));
    }

    @PostMapping("/user")
    public Result<DABEUser> createUser(String fileName, String userName){
        return handleUser(dabeService.createUser(fileName, userName));
    }

    @PostMapping("/user/attr")
    public Result<DABEUser> declareAttr(String fileName, String attrName){
        return handleUser(dabeService.declareAttr(fileName, attrName));
    }

    private Result<DABEUser> handleUser(DABEUser user) {
        if (user == null) {
            return Result.internalError("no user");
        } else {
            return Result.okWithData(user);
        }
    }

    @GetMapping("/user/apply")
    public Result<Object> approveAttrApply(String fileName, String attrName, String toUserName){
        return dabeService.approveAttrApply(fileName, attrName, toUserName)
            .getResult(str -> str);
    }
}
