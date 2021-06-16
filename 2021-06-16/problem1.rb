require 'set'
require 'benchmark'

class ClassName
    def initialize
        
    end
    def makearray(min,max,size)
        return Array.new(size){rand(min..max)}
    end

    def methodonsq(arr,k)
        arr.each do |x|
            arr.each do |y|
                if x+y ==k
                    return true 
                end
            end 
        end
        return false
    end   

    def methodon(arr,k)
        seen = Set.new
        arr.each do |x|
            if seen === k- x
                return true
            end
            seen.add(x)        
        end
        return false
    end
end

a = ClassName.new
arr = a.makearray(10,10000000)
puts arr.length

time = Benchmark.measure{
    puts a.methodon(arr,21)
    
}
#puts a.methodonsq(arr,21)
puts time.real


