class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        
        d = defaultdict(list)
        
        for s in strs:
            current_count = [0] * 26
            for l in s:
                i = ord(l) - ord('a')
                current_count[i] += 1
            d[tuple(current_count)].append(s)
        return d.values()