package com.weiyan.atp.service;

import com.weiyan.atp.data.request.web.ShareContentRequest;
import com.weiyan.atp.data.response.intergration.EncryptionResponse;
import com.weiyan.atp.data.response.web.PlatContentsResponse;

import javax.validation.constraints.NotEmpty;

/**
 * @author : 魏延thor
 * @since : 2020/6/11
 */
public interface ContentService {
    void shareContent(ShareContentRequest request);

    String decryptContent(@NotEmpty String cipher, @NotEmpty String fileName);

    /**
     * tag和fromUserName不能同事为空
     */
    PlatContentsResponse queryPlatContents(String fromUserName, String tag,
                                           int pageSize, String bookmark);

    EncryptionResponse encContent(ShareContentRequest request);
}
