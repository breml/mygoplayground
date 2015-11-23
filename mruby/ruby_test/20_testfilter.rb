require './myconfig'
class TestFilter
        include MyConfig

        config :message, :default => "defaultvalue"

        public
        def initialize()
                @message = self.class.get_config["message"][:default]
        end

        public
        def filter(event)
                event["message"] = @message
                event
        end
end