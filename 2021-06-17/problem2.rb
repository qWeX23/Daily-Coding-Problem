def arrayProduct(arr)
    productArr = Array.new(arr.size)

    product = multiplyArr(arr)

    arr.each_with_index do |x,i|
        productArr[i]=product/x
    end 
    return productArr
end

def multiplyArr(arr)
    product = 1
    arr.each do |x|
        product = x* product
    end
    return product
end


def mathArr(arr,idx)#array, index to exclude
    product = 1
    arr.each_with_index do |x,i|
        if idx !=i 
            product= product*x
        end
    end
    return product
end
arr = [1,2,3,4,5]


#Array.new(4){rand(0..4)}

puts arrayProduct(arr)