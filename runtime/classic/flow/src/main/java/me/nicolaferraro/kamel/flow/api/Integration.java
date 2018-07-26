package me.nicolaferraro.kamel.flow.api;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Integration {

    /**
     * This is the only relevant field
     */
    private IntegrationSpec spec;

    public Integration() {
    }

    public IntegrationSpec getSpec() {
        return spec;
    }

    public void setSpec(IntegrationSpec spec) {
        this.spec = spec;
    }

}
