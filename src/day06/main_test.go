package main

import (
	_ "embed"
	"testing"
)

func TestDay06a(t *testing.T) {
	inputs := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    7,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
		"nppdvjthqldpwncqszvftbrmjlhg":      6,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
		"pjbjvjtjljplppjssvtvwtwptptztltbtrrjgrjrzrqrjrbrhbrhrlllbpbdbbzqqgsqshqssjjbsjbsbmbhhmchhrqrcqqbwbqwwqrrznnsbswwwdjwdwmmsvszzlbbgddbgdgfgttzjzrjzrjzzvrvqqgpqggtbgtgvvrhvhtvtjjbpjjfjhjbhbddbjjjmzmtmgmpgmmmljlnljjmpmbpmbmrrlhhlppdgdfgfbbqlbbtffjjgvvnpvpbpttqmmnhhgfgrrwhrrbnbznzccmbmvmmzszsvsbvbccsrslsjsbbtdtwtvttvpvzzvbzzwczwcczhzhwhjhjghgppgpgttdwdhwhphnppqmpqqhthcchdhmhnnbcnbcbbggbfblljttwsswspsggpjjzszcscsmcmdmnngzzhrzrbrhbhzbzvzgvgffnlnljlrrhchhsvhssmpmccncdnccgdgbglgtlggnllsvvpfvpfpbffsgglddjrrzzphhptprtppwrwffzllrbrmrpmpdmpdmpplpspcphcphhgmmqnmmvnmmdvmvsmmqjmmlmlbmlbmlltlptphpnncscbscbcggwhhgjgqjqmjmdjmdjdrrvgvvzlldgdnnvttmmpffdjjvvchcwwbhwbwzzlmlccrttcntccpcgccpgcpggdrrbtrthhlrrbqrrpspdsdldbldblbzbpbgggtngtgqqtwwdjjmmcrmcrcvclcddhllpzpdzzmccrtccfvffccfhccpscctbbbzqzvvllgwlgwgvwwjswjswsvvwhhvjvsjvjmmjhjrhrmrvrnrccmnmzzmdzzbtbvbqvvgzgcgvvvvlltvllbfbqqrppwhpwpffzddzdzwdzzrggmhmfhmhbbzjzsjzzhhjdhhdnndsnssnfffbmffwhwrhhmmbnnbrnrbrrtqtztnzzzzblbhbdhbbfmfqfmmsgszzvfzfmfwfnwfwggwngnqnwqnwqwvwqvwvdvrvjjfnjnmnfmmwzzltztjjqnnnmlmzlmmrcmclcqqhrhdhccdfdvfvccvtcvvdmmtmccwjjcbcrrjmrjrdrffgwwvbvlvsspwpsspzsschssmqsqmqmlqmmqgqfqcqjcjtthjttlddfvvwwjvvtpvvfsvffqnffznzqzszgzmmjttwztwzzhqqccqsqmqnmqqjhqhzhwhvhdddsndsdfftvffwlwnnmmdpmmnhhrqrrclcdlcddhcdcppgrprnpnptnpphgpbqfngdgzvgndwcgrwcsfmhzsvddhzbgjmvvdjjzswvgnpmvgdpwsgbgjzjpsrfdzdzjzzrpplbhsmgddqzjbdzdzltqqwqjzqvwfmcdppbdbprrwzhmnrqclzrnmdjnfbwmvdrwtpwvgscrqgpndqnzbjsbljcbthbpgdjdcdwfhpvjnbsfjdlrjldvvmtfdslrhlfwmvclqrljrqmmjgqfwmfgwdjzzptgcthvtgdswsqjrqvnzmtqldjjcqnfhtvbwhjqlvpptfwjrdpcwvzddgcjzvqbhtsnnnjqqmqlbgvqmvjhvvpbzcbdmhgmcjbfcccsvlzjztvjzrrlhtgwccdcgcptqlmdhmdhvqzfntbjqtsmvqgwsltqntgszllntrljfgfsghtbbcqrdgwqphmbqtzmjqccrgvqpqpchzjstdmmtvntwjqsbcqjgnhzlllcfbpgtgrhwwhqqdlgrlsbzbmchvjnsgpdnmqvtgwqjpgflqgfngjfcfwqzmvvgzmmhbgfnbzvzclwclqdcccgbrrzpwdtprgsvhbgsnbntgrvnzhrnzfzdmnlbnrbqvmjbwpgvjlhbcvsrlqmcsnlrvtfdwtvcbmlndgbctsnmtctjszlpddqmzbtphhhfznwbdfsgppmdmczmhmmrzpllfqqbgvlsrscpfgznhdhgrnnnvrchgvzlqbgvcfghjvlvrvpclfcshbmvglcfrjbzrbcjmjjrfgqthwfrqbgtjldmbnfwllspmwrvstvrltvrlvrtjvprgtgzjlrgclvjhqpfcwcdbdtzwdsdfrtsvtvgjmsszdfqlmhqqlzswjfndswlmhcrhglphvpnfjpbmggbwlmzjchpnrllbjpmgmzjjrqpqgsbrszqhdljcpnclvrvbntgtcdcmhtdhgslhpvdjpvrszfrjhsbvcvtfwvvgczprnpbhmnnlmctbtqdjspgvhvnhwvspwgnjvzllwlnjhfjwsslppmjbfbdnthcpzbcmnnbvhctgwgdvhvlrbltmdnlfcsncqgrmjprshdvvtvcccgzhszcjgczhmhtvmccjpchqshhdzjjhbfpzqdjszdhdvlmgctmwcjprwlsqbcqhlcrfdgnqzfdfvqslmqlppbsvbmjmfbrtdmpmtqvwvppcfzddjzhhzlrrnnhbrlhmzlqwftprfvctnfhfhfzrnrvggfqmqwcwszhtbfjncprgwcqbjlvtnrprlwwghswvprjmsbmqvwnnfggprndvshfvvwtrqjpwghgbppftgzhqjslfzhngwfsjnmjzdsjqgpmglwnjlcgmczgvndszrszcpnzqpbzjmgrfsbjlghwrbqsqdhlnhzsvsgbqhbcdffjlgrbdrrjvclzqpftlhdvvcrvlgvlpnjcqdcbdjtlwnldjhhhzrwsqlhlsztwrznfsszptlrhjmqwmnfwjtjwmmmtvwzhpmjgzgsscbddgvvhpcnhvnggzhbzvvjlmdftpbcsvtsttrvgghptmmcdclbdvmnsdntthfbdznbclwccnlzcvdwzrqgddjszvbdqcjppzrtpnrhfcvvwpjqczgqwzzzvzmlnlzqszvtllftthgwgftjzsndpzzcnqpcvmsdvvfrjdsvfclsqqhsjrrctfvdrlhfmhprjggdcmqrrbqtwnrllhhztvjgmzqszbvqfwsgllvhsvfrjffvdscwjzqlzlwdpgthddpgzjfdbdqpsnntwpslvsdpqfnsgcllszcjwvtqhwhpfrlfdgwrfmgfpjmvnstrmtfcvgwlqdfqvntltqtrmjjtwcthvwntqgvncssplnmvlnstlcphvlcmvjnstwldtntchcbmzmlzhgjfbrdlgzvqpgcndmfdnmcnwhmpdnpqstfddddcrpgrpfwfbzjqtnzwwqpzrqpmrjpfznrndfgwhtlvrcrphqfjzjbttwhgnsngqwvnsbvcqtjlmhvmnmnnmjcmlpnpgmrqsbmgljvsfqvrlljqzmzqqbgpvcrwdjmgsglssjswmnvtshhfqjhqmfmvcjwfpwsppgtrqsbhhcdljnjphnjszqpvdplbwzpwmmpwfhmhngtllzqvpmgdctmfqwwqjszssmjhwnrjdtmmvpdnwlqtcbpfcmwtbjmmsmmdpqgzdhsblgjmjbpzgqvqhnggtwmhztbbhlflllgwblncjjsngdgvsfdmsbsvlpnjjzqqbzhsqclmjnnmmwlpvtgwqmcgmrqdwdddlgbvhntbztbjnqhdlggnzwsdtdzprgddhtcttjrcpszgchtfwqjsdlnbntfwqpzpfsqrqjhthmcfszwtwcqwbvfzdnrrpmzjdrhsgmhfbsldvcrjdwvpqpszzlvbptljgvccqsdhhnztjpghbvhfptgplqdvldjzfthpspwvgljwnnndwrqzbrstnqbvrrcghssnrpvtrhmvcmbngwndzfswmgjwnnzqdcjhpthcgvthsnwqzrnzrvdjmctchhsbnrtvctzqfpcjhzmhnfjlqftbjztfbcppgmwvrzzrvlcpnpwwpvtcpdplrcfpgfqjtlfjtphhpcltwqcbqbznbtjrtdrpgtvzmgsclhpptrssqqbctdrftqzmwjmrmjtgmjmsnbnspjvcqpqnmgzgjrmfhghvsfsdqnbdjsbcpczsdswdcvhfzlgpzbtmztcnbpcvjnlcdmmlbtwzsfqtfnlrwjtwmgslcgptgbdsfwdhppvfwbbgdfdqtrbncbznmqtchzsdzlhlhjnnbpdvnnfjrdfbdqmvcb": 1275,
	}
	for k, v := range inputs {
		actual := day06(k, 4)
		expected := v
		if actual != expected {
			t.Fatalf(`actual = %v, expected = %v`, actual, expected)
		}
	}
}

func TestDay06b(t *testing.T) {
	inputs := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
		"nppdvjthqldpwncqszvftbrmjlhg":      23,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
		"pjbjvjtjljplppjssvtvwtwptptztltbtrrjgrjrzrqrjrbrhbrhrlllbpbdbbzqqgsqshqssjjbsjbsbmbhhmchhrqrcqqbwbqwwqrrznnsbswwwdjwdwmmsvszzlbbgddbgdgfgttzjzrjzrjzzvrvqqgpqggtbgtgvvrhvhtvtjjbpjjfjhjbhbddbjjjmzmtmgmpgmmmljlnljjmpmbpmbmrrlhhlppdgdfgfbbqlbbtffjjgvvnpvpbpttqmmnhhgfgrrwhrrbnbznzccmbmvmmzszsvsbvbccsrslsjsbbtdtwtvttvpvzzvbzzwczwcczhzhwhjhjghgppgpgttdwdhwhphnppqmpqqhthcchdhmhnnbcnbcbbggbfblljttwsswspsggpjjzszcscsmcmdmnngzzhrzrbrhbhzbzvzgvgffnlnljlrrhchhsvhssmpmccncdnccgdgbglgtlggnllsvvpfvpfpbffsgglddjrrzzphhptprtppwrwffzllrbrmrpmpdmpdmpplpspcphcphhgmmqnmmvnmmdvmvsmmqjmmlmlbmlbmlltlptphpnncscbscbcggwhhgjgqjqmjmdjmdjdrrvgvvzlldgdnnvttmmpffdjjvvchcwwbhwbwzzlmlccrttcntccpcgccpgcpggdrrbtrthhlrrbqrrpspdsdldbldblbzbpbgggtngtgqqtwwdjjmmcrmcrcvclcddhllpzpdzzmccrtccfvffccfhccpscctbbbzqzvvllgwlgwgvwwjswjswsvvwhhvjvsjvjmmjhjrhrmrvrnrccmnmzzmdzzbtbvbqvvgzgcgvvvvlltvllbfbqqrppwhpwpffzddzdzwdzzrggmhmfhmhbbzjzsjzzhhjdhhdnndsnssnfffbmffwhwrhhmmbnnbrnrbrrtqtztnzzzzblbhbdhbbfmfqfmmsgszzvfzfmfwfnwfwggwngnqnwqnwqwvwqvwvdvrvjjfnjnmnfmmwzzltztjjqnnnmlmzlmmrcmclcqqhrhdhccdfdvfvccvtcvvdmmtmccwjjcbcrrjmrjrdrffgwwvbvlvsspwpsspzsschssmqsqmqmlqmmqgqfqcqjcjtthjttlddfvvwwjvvtpvvfsvffqnffznzqzszgzmmjttwztwzzhqqccqsqmqnmqqjhqhzhwhvhdddsndsdfftvffwlwnnmmdpmmnhhrqrrclcdlcddhcdcppgrprnpnptnpphgpbqfngdgzvgndwcgrwcsfmhzsvddhzbgjmvvdjjzswvgnpmvgdpwsgbgjzjpsrfdzdzjzzrpplbhsmgddqzjbdzdzltqqwqjzqvwfmcdppbdbprrwzhmnrqclzrnmdjnfbwmvdrwtpwvgscrqgpndqnzbjsbljcbthbpgdjdcdwfhpvjnbsfjdlrjldvvmtfdslrhlfwmvclqrljrqmmjgqfwmfgwdjzzptgcthvtgdswsqjrqvnzmtqldjjcqnfhtvbwhjqlvpptfwjrdpcwvzddgcjzvqbhtsnnnjqqmqlbgvqmvjhvvpbzcbdmhgmcjbfcccsvlzjztvjzrrlhtgwccdcgcptqlmdhmdhvqzfntbjqtsmvqgwsltqntgszllntrljfgfsghtbbcqrdgwqphmbqtzmjqccrgvqpqpchzjstdmmtvntwjqsbcqjgnhzlllcfbpgtgrhwwhqqdlgrlsbzbmchvjnsgpdnmqvtgwqjpgflqgfngjfcfwqzmvvgzmmhbgfnbzvzclwclqdcccgbrrzpwdtprgsvhbgsnbntgrvnzhrnzfzdmnlbnrbqvmjbwpgvjlhbcvsrlqmcsnlrvtfdwtvcbmlndgbctsnmtctjszlpddqmzbtphhhfznwbdfsgppmdmczmhmmrzpllfqqbgvlsrscpfgznhdhgrnnnvrchgvzlqbgvcfghjvlvrvpclfcshbmvglcfrjbzrbcjmjjrfgqthwfrqbgtjldmbnfwllspmwrvstvrltvrlvrtjvprgtgzjlrgclvjhqpfcwcdbdtzwdsdfrtsvtvgjmsszdfqlmhqqlzswjfndswlmhcrhglphvpnfjpbmggbwlmzjchpnrllbjpmgmzjjrqpqgsbrszqhdljcpnclvrvbntgtcdcmhtdhgslhpvdjpvrszfrjhsbvcvtfwvvgczprnpbhmnnlmctbtqdjspgvhvnhwvspwgnjvzllwlnjhfjwsslppmjbfbdnthcpzbcmnnbvhctgwgdvhvlrbltmdnlfcsncqgrmjprshdvvtvcccgzhszcjgczhmhtvmccjpchqshhdzjjhbfpzqdjszdhdvlmgctmwcjprwlsqbcqhlcrfdgnqzfdfvqslmqlppbsvbmjmfbrtdmpmtqvwvppcfzddjzhhzlrrnnhbrlhmzlqwftprfvctnfhfhfzrnrvggfqmqwcwszhtbfjncprgwcqbjlvtnrprlwwghswvprjmsbmqvwnnfggprndvshfvvwtrqjpwghgbppftgzhqjslfzhngwfsjnmjzdsjqgpmglwnjlcgmczgvndszrszcpnzqpbzjmgrfsbjlghwrbqsqdhlnhzsvsgbqhbcdffjlgrbdrrjvclzqpftlhdvvcrvlgvlpnjcqdcbdjtlwnldjhhhzrwsqlhlsztwrznfsszptlrhjmqwmnfwjtjwmmmtvwzhpmjgzgsscbddgvvhpcnhvnggzhbzvvjlmdftpbcsvtsttrvgghptmmcdclbdvmnsdntthfbdznbclwccnlzcvdwzrqgddjszvbdqcjppzrtpnrhfcvvwpjqczgqwzzzvzmlnlzqszvtllftthgwgftjzsndpzzcnqpcvmsdvvfrjdsvfclsqqhsjrrctfvdrlhfmhprjggdcmqrrbqtwnrllhhztvjgmzqszbvqfwsgllvhsvfrjffvdscwjzqlzlwdpgthddpgzjfdbdqpsnntwpslvsdpqfnsgcllszcjwvtqhwhpfrlfdgwrfmgfpjmvnstrmtfcvgwlqdfqvntltqtrmjjtwcthvwntqgvncssplnmvlnstlcphvlcmvjnstwldtntchcbmzmlzhgjfbrdlgzvqpgcndmfdnmcnwhmpdnpqstfddddcrpgrpfwfbzjqtnzwwqpzrqpmrjpfznrndfgwhtlvrcrphqfjzjbttwhgnsngqwvnsbvcqtjlmhvmnmnnmjcmlpnpgmrqsbmgljvsfqvrlljqzmzqqbgpvcrwdjmgsglssjswmnvtshhfqjhqmfmvcjwfpwsppgtrqsbhhcdljnjphnjszqpvdplbwzpwmmpwfhmhngtllzqvpmgdctmfqwwqjszssmjhwnrjdtmmvpdnwlqtcbpfcmwtbjmmsmmdpqgzdhsblgjmjbpzgqvqhnggtwmhztbbhlflllgwblncjjsngdgvsfdmsbsvlpnjjzqqbzhsqclmjnnmmwlpvtgwqmcgmrqdwdddlgbvhntbztbjnqhdlggnzwsdtdzprgddhtcttjrcpszgchtfwqjsdlnbntfwqpzpfsqrqjhthmcfszwtwcqwbvfzdnrrpmzjdrhsgmhfbsldvcrjdwvpqpszzlvbptljgvccqsdhhnztjpghbvhfptgplqdvldjzfthpspwvgljwnnndwrqzbrstnqbvrrcghssnrpvtrhmvcmbngwndzfswmgjwnnzqdcjhpthcgvthsnwqzrnzrvdjmctchhsbnrtvctzqfpcjhzmhnfjlqftbjztfbcppgmwvrzzrvlcpnpwwpvtcpdplrcfpgfqjtlfjtphhpcltwqcbqbznbtjrtdrpgtvzmgsclhpptrssqqbctdrftqzmwjmrmjtgmjmsnbnspjvcqpqnmgzgjrmfhghvsfsdqnbdjsbcpczsdswdcvhfzlgpzbtmztcnbpcvjnlcdmmlbtwzsfqtfnlrwjtwmgslcgptgbdsfwdhppvfwbbgdfdqtrbncbznmqtchzsdzlhlhjnnbpdvnnfjrdfbdqmvcb": 3605,
	}
	for k, v := range inputs {
		actual := day06(k, 14)
		expected := v
		if actual != expected {
			t.Fatalf(`actual = %v, expected = %v`, actual, expected)
		}
	}
}
