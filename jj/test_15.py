def solution(S, T):
    # 计算最少操作数
    min_ops = len(S)  # 初始化为S的长度（最大删除次数）
    
    # 循环长度应该是S和T的最小长度
    min_len = min(len(S), len(T))
    
    # 尝试让S成为T的每个可能前缀
    for prefix_len in range(min_len + 1):
        ops = 0
        
        # 如果S的长度大于当前前缀长度，需要删除多余的字符
        if len(S) > prefix_len:
            ops += len(S) - prefix_len
        
        # 计算需要修改的字符数（比较S和T的前prefix_len个字符）
        for i in range(prefix_len):
            if S[i] != T[i]:
                ops += 1
        
        if ops < min_ops:
            min_ops = ops
    
    return min_ops

def solution_with_debug(S, T):
    print(f"S='{S}', T='{T}'")
    
    min_ops = len(S)  # 初始化为S的长度（最大删除次数）
    
    # 循环长度应该是S和T的最小长度
    min_len = min(len(S), len(T))
    
    for prefix_len in range(min_len + 1):
        ops = 0
        prefix = T[:prefix_len]
        
        # 删除操作
        if len(S) > prefix_len:
            delete_ops = len(S) - prefix_len
            ops += delete_ops
            print(f"  前缀'{prefix}': 删除{delete_ops}个字符", end="")
        
        # 修改操作
        modify_ops = 0
        for i in range(prefix_len):
            if S[i] != T[i]:
                modify_ops += 1
        ops += modify_ops
        
        if modify_ops > 0:
            print(f", 修改{modify_ops}个字符", end="")
        
        print(f" -> 总操作数: {ops}")
        
        if ops < min_ops:
            min_ops = ops
    
    print(f"最少操作数: {min_ops}\n")
    return min_ops

# 测试用例
print("Testing solution function with debug info:")
solution_with_debug('aba', 'abb')
solution_with_debug('abcd', 'efg')
solution_with_debug('xyz', 'xy')
solution_with_debug('hello', 'helloworld')
solution_with_debug('same', 'same')

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