class Stairs
    def initialize
        
    end
    def getStairs(arr,start=0,bigArr=[[]])
        steps_arr = []
        step_size = 2
        if start >= arr.length
            return bigArr
        end

        arr.each_with_index do |x,i|
            
            next if x=! 1 #already has been stepped
            if i<start #need to get past this level
                steps_arr.append(x)
            end
            next if arr.length <= i+step_size-1 #no more steps to take

            #is this a valid step?
            # (i..i+step_size).each do |j|
                
            # end

            #add step and fill to end 
            steps_arr.append(step_size)
            fill1s = arr.length-(i+step_size-1)
            steps_arr.append(Array.new(fill1s-1){1})
           
         end
         bigArr.append(steps_arr)
         puts(steps_arr.length)
         getStairs(steps_arr,start+1,bigArr)
    end
end

s = Stairs.new

arr = Array.new(3){1}
puts s.getStairs(arr)


