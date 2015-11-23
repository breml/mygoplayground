require './testfilter'

tf = TestFilter.new
puts tf.filter({"message" => "input", "key" => "value"})