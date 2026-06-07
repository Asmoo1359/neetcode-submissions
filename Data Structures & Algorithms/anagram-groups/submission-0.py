class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        myMap = {}
        myList = []
        for s in strs :
            Temp = sorted(s)
            stringTemp = ''.join(Temp)
            if stringTemp in myMap :
                myMap[stringTemp].append(s)
            else :
                myMap[stringTemp] = [s]

        for item in myMap :
            myList.append(myMap[item])
        return myList
