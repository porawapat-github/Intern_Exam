def wall():
    Wall1 = [1,2,3,4,5,6]
    Wall2 = [1,2,2,1,5,6]
    return Wall1, Wall2

def bricks_added(wall_heights):
    total_bricks = 0
    last_brick_height = 0

    print("\nการคำนวณทีละขั้น:")
    for i, height in enumerate(wall_heights):
        if height > last_brick_height:
            bricks_needed = height - last_brick_height
            total_bricks += bricks_needed
            print(f"ขั้นที่ {i}: {last_brick_height} -> {height} = +{bricks_needed} อิฐ")
        else:
            print(f"ขั้นที่ {i}: {last_brick_height} -> {height} = +0 อิฐ (ไม่ต้องเพิ่ม)")
        last_brick_height = height

    return total_bricks

def main():
    wall1, wall2 = wall()
    
    # Case 1
    print("Case 1:")
    print(f"Input: {wall1}")
    result1 = bricks_added(wall1)
    print(f"รวมอิฐที่ต้องใช้: {result1}\n")
    
    # Case 2
    print("Case 2:")
    print(f"Input: {wall2}")
    result2 = bricks_added(wall2)
    print(f"รวมอิฐที่ต้องใช้: {result2}")

if __name__ == "__main__":
    main()