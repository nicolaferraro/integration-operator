package me.nicolaferraro.kamel.flow;

import org.apache.camel.RoutesBuilder;
import org.apache.camel.component.mock.MockEndpoint;
import org.apache.camel.test.junit4.CamelTestSupport;
import org.junit.Test;

import java.io.InputStreamReader;
import java.io.Reader;

public class YamlRouteBuilderTest extends CamelTestSupport {

    @Test
    public void testParser() throws Exception {
        MockEndpoint ciao = getMockEndpoint("mock:ciao");
        ciao.expectedMinimumMessageCount(5);
        ciao.assertIsSatisfied();
    }

    @Override
    protected RoutesBuilder createRouteBuilder() throws Exception {
        try (Reader in = new InputStreamReader(getClass().getResourceAsStream("/example-1.yaml"))) {
            return new YamlRouteBuilder(in);
        }
    }
}
