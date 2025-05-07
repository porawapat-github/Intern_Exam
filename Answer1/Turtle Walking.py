# ฟังก์ชันข้อมูลทางเดินของเต่าทั้งหมด
def turtle():
    Turtle1 = [
        [7,2,0,1,0,2,9],
        [8,4,8,6,9,3,3],
        [7,8,8,8,9,0,6],
        [4,7,2,7,0,0,7],
        [6,5,7,8,0,7,2],
        [8,1,8,5,4,5,2]
    ]

    Turtle2 = [
        [7,2,0],
        [8,4,8],
        [7,8,8]
    ]

    Turtle3 = [
        [7,2,0],
        [8,4,8]
    ]

    return Turtle1, Turtle2, Turtle3

# ฟังก์ชันที่เต่าจะเดิน
def caseall(data, name="Turtle"):
    rows = len(data)
    cols = len(data[0])
    result = []

    for col in range(cols):
        if col % 2 == 0:
            for row in range(rows):
                result.append(data[row][col])
        else:
            for row in reversed(range(rows)):
                result.append(data[row][col])

    print(f"\nผลลัพธ์ที่เต่าเดินใน {name}:")
    print(*result)

# เรียกใช้ฟังก์ชันหลัก และแสดงผลทั้งหมด
Turtle1, Turtle2, Turtle3 = turtle()


print("Turtle1:")
for row in Turtle1:
    print(row)
caseall(Turtle1, "Turtle1")


print("\nTurtle2:")
for row in Turtle2:
    print(row)
caseall(Turtle2, "Turtle2")


print("\nTurtle3:")
for row in Turtle3:
    print(row)
caseall(Turtle3, "Turtle3")
