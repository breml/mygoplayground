module MyConfig
        module DSL
                def config(name, opts={})
                        @config ||= Hash.new
                        name = name.to_s if name.is_a?(Symbol)
                        @config[name] = opts
                        if name.is_a?(String)
                                define_method(name) { instance_variable_get("@#{name}") }
                                define_method("#{name}=") { |v| instance_variable_set("@#{name}", v) }
                        end
                end
                def get_config
                        return @config
                end
        end


        def self.included(base)
                # Add the DSL methods to the 'base' given.
                base.extend(MyConfig::DSL)
        end
end