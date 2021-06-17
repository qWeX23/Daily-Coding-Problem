def stairs(n)
    if n<=1 
        return 1
    end
    return stairs(n-1) + stairs(n-2)
end

def stairs2(n)
    a=1
    b=2
    c=0
    (1...n).each do |x|
        c=b
        b = b+a
        a=c
        puts "x: %d a: %d b: %d"%[x,a,b]

    end
    return(a)
end

puts stairs2(5)