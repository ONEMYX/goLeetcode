def solution(S, T):
    # 计算最少操作数
    min_ops = len(S) + len(T)  # 初始化为一个很大的值
    
    # 尝试让S成为T的每个可能前缀
    for prefix_len in range(len(T) + 1):
        ops = 0
        
        # 如果S的长度大于当前前缀长度，需要删除多余的字符
        if len(S) > prefix_len:
            ops += len(S) - prefix_len
        
        # 计算需要修改的字符数（比较S和T的前prefix_len个字符）
        for i in range(min(prefix_len, len(S))):
            if S[i] != T[i]:
                ops += 1
        
        # 如果S的长度小于前缀长度，需要添加字符（通过修改操作）
        if len(S) < prefix_len:
            ops += prefix_len - len(S)
        
        if ops < min_ops:
            min_ops = ops
    
    return min_ops

# 测试用例
print("Testing solution function:")
print(f"solution('aba', 'abb') = {solution('aba', 'abb')} (expected: 1)")
print(f"solution('abcd', 'efg') = {solution('abcd', 'efg')} (expected: 4)")
print(f"solution('xyz', 'xy') = {solution('xyz', 'xy')} (expected: 1)")
print(f"solution('hello', 'helloworld') = {solution('hello', 'helloworld')} (expected: 0)")
print(f"solution('same', 'same') = {solution('same', 'same')} (expected: 0)")

# 验证所有测试用例
assert solution('aba', 'abb') == 1
assert solution('abcd', 'efg') == 4
assert solution('xyz', 'xy') == 1
assert solution('hello', 'helloworld') == 0
assert solution('same', 'same') == 0

print("\nAll tests passed!") 