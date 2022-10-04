package com.weiyan.atp.data.request.web;

import lombok.Data;

import java.util.List;

import javax.validation.constraints.NotEmpty;
import javax.validation.constraints.NotNull;

/**
 * @author : 魏延thor
 * @since : 2020/6/1
 */
@Data
public class ShareContentRequest {
    @NotEmpty
    private String fileName;

    @NotNull
    private List<String> tags;

    @NotEmpty
    private String plainContent;

    /**
     * (A AND B AND (C OR D))
     */
    @NotEmpty
    private String policy;
}
