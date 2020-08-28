package arrayandmatrix

import (
	"testing"
)

// 给定一组唯一的单词， 找出所有不同 的索引对(i, j)，使得列表中的两个单词， words[i] + words[j] ，可拼接成回文串
// 336. 回文对
// https://leetcode-cn.com/problems/palindrome-pairs/
func Test_PalindromePairs(t *testing.T) {
	words := []string{"ha", "eaebjghbjjheddgfa", "hhaccedbbdbciid", "ediihegjbefacccga", "hdbiiicacdcdf", "cefafdjbafbiahcc", "abhdedaagbbcibhbjbbd", "dcbefdcfceaiaahfejc", "jbhgfdcfihg", "g", "dbi", "hgabcfibbiddcji", "hjhdiiiaicij", "afbbcbdffeifafc", "abbihgiiif", "be", "j", "ecjefbbcjac", "eg", "ighcicacd", "eicceeaaefg", "ia", "fcfgjaijgdbchhfhddg", "aghecechf", "ihgcgafga", "bjfgdjba", "ab", "hhbjhe", "hhedaebbceeeja", "eabgicejiddiijjdbjjj", "icb", "eghhbhfajccb", "fideijjhheabhjfhh", "def", "bihfa", "hibhifghhddcijfcfebc", "hihdgbiijh", "jcibccjei", "bgdcdcbcjaicihjiafg", "jiie", "bbihfgcffheccdeddd", "hefea", "iagegahcehdcf", "fghhgcfigegcfbcde", "igbch", "cicdjechehfdiagcc", "djcieegaehigc", "hgajaehbdg", "ejbcdehfibdecgbhb", "bcbcfdgefffi", "aabjeadagfifeihjac", "hjhdgcaddbcf", "idiggbfihfiic", "gcffh", "e", "cfggajjfj", "gegedie", "achiajcjecja", "bbbcgjefbhhf", "iadjhigibgbh", "fgdbicdg", "jghfhhcahah", "hajjfaeahbhiijegab", "fafhba", "idhgehaiafhibe", "gjfcgjbejhjgfbg", "ajai", "gbhgefeicjeejhghcb", "i", "di", "hbcahahdfbdfcbffi", "hedigibfiacebff", "jcddfcjjagbbgh", "dchca", "hgagfghbbhb", "jidi", "cgcjcdhcejj", "chhdbhejijgaehbaic", "gabedga", "fjdbegjfjaeagei", "ifefffgd", "jghhhicgijjdeejc", "idadhej", "hgcgagfhhbj", "bhhhgbihga", "hgdjdehfi", "fbaciggbdg", "jdbj", "dhgeafc", "fjjjj", "djghaahiji", "bhegiaiicgijfjacgacj", "eaheeg", "hjdcedga", "dahbieb", "ejgahijifchajajee", "haejbc", "ebijhbf", "fed", "ddibabcdhdjdgeie", "ecg", "ihebibfhdifagdcdbjdi", "cghijahbedjccg", "egiaj", "egdhcbcihb", "adhbbdgcffcgggehgbi", "eaiii", "gagejb", "dfhccaadifhjjdfffiba", "bdhfhahdbefgifgdf", "eibhdehbfccbc", "jgbcgbebbjf", "feagecfhiafh", "ifagegfbaegefihifja", "iicgcchcie", "defjbi", "djdejegjefffbggcgf", "dejjggighbcgidjgfbc", "ebibeejhjchjadahcd", "ccidbehccbhfhbdj", "ibdbecfgicajdg", "eegjbgceifhfahjiefe", "bgfdddajeacjfddhheie", "egghhabi", "dafaidibhdebijhcji", "igj", "eihif", "hjjdcjcihi", "ihffbe", "bf", "ajdfbe", "jihcc", "cgejejhbfffjbca", "cajgii", "abceahigj", "hhbhehfbhfefajbiceia", "fgghdgfbbeicj", "ghhaibheccdgdabeaa", "ahidijigdfdbcefai", "jgahjjhdejc", "aabhjdfbeghffgiihhec", "fhdgfeiijga", "ffgjeiihgdhbidehgci", "jdcaaafe", "dbabfbfaheidce", "bcfeeihjdjgfh", "ifigghggicj", "biibie", "jajigfecdaagfeggjeh", "igidegf", "bffhefffjahgjbdijh", "egcbfjgc", "gbghhgdbb", "bghbijfeij", "hagfbg", "iiigdha", "cfdchgbabijjbidfje", "ggeee", "hidfaajfb", "cjab", "hgff", "daahgg", "cf", "fjhfaiheeicdib", "giegdhbajcgagf", "bbebjef", "f", "jhhiaebdcec", "fdjjjgaeeg", "ccbjaageiaddbc", "fijdjbijea", "jggj", "ieaaeiiddiaj", "ghg", "deiaicejhggcfchj", "hhhef", "fcecgfha", "hgejaji", "hce", "abfhgjegiefjb", "iijibjdheejaieiga", "jgbifgcjedi", "jfafebjfjadcibcbjdh", "ai", "hihjhcachbcghacib", "cicceaebdidgbef", "hecfiagefhhcjcbdhicd", "diafccjdb", "dffej", "giadddbjjigghbjbdfib", "ahd", "efafdc", "cgejjgjijh", "cjbf", "aceahd", "hcajghbdibfhb", "aacabbjhfjcge", "fhggbacdiaed", "bahehcjeifhcghdhi", "fbhdbajbcjdhe", "jbahjdiiffb", "dgechbigjabghchfbei", "ddbbbb", "ebiii", "jjbbgdcgdebijiejc", "jhjef", "ggbdgciibiif", "gice", "ffjccagjigbd", "jbefedejfcafhe", "dfd", "gbibjaachijbghhffbi", "fdddefbdjh", "aadjbjdfedc", "hbiecdjjfci", "jci", "jhj", "aejeieccbfi", "deid", "giggagjhcbgbbe", "jgfjfa", "jibhffhbae", "hd", "iedidfdachgcifbjag", "fjj", "hbbddajidfjiidd", "fgjggedajihdhjd", "ijgjbfjf", "gbjgdceeaej", "bgaijh", "egcgbegfadeibcfcebf", "bccaedehchiaed", "jgaebdedhbi", "hjafcghabb", "iiid", "aiaefejedcjhibbae", "fd", "bieciafejaa", "ihf", "aaaijicifa", "gegjgcceedi", "dbfbcggfdfa", "cdjb", "a", "ebcheedccdaaidjfi", "cchiaaagfghdhefjfiib", "ecdjca", "bdfgjcgcheebdggide", "jfigaciejefba", "idjgdbffch", "aaedgfgjhcbigbccgji", "gfehgjchgbejea", "dagbhdija", "bdhcfahifjbigj", "cifaeicajhfdgajjfj", "bhigd", "icbbacfejcddj", "dc", "cebgi", "cbicfejcjifiihd", "eeeic", "fbighdhg", "hjefedih", "b", "jheheehjcbhjadecb", "jeeidjdaiagcg", "dadhghdecedahbcicg", "edije", "jacdjiadbibejjccgd", "jdbdejehagjjfc", "fdcgbb", "fgihbdjchfgbjbafcfj", "dfbjfbidafegcfcgdcgf", "hiccidgdbbbbgijcf", "eagichhf", "ifhjbgheiaaagde", "bfggb", "bdfcjeihfdahec", "bjhbbedgaaacg", "eeaihbj", "ihfacji", "dajjaajdgjdjc", "jhibebgifihi", "jfaadbaeiiciehihecc", "hfhge", "bab", "eaecdfc", "dhejdcejhfbabddbb", "gfaiffdacigjeidead", "debbibeeajdf", "dfdbahhiigjiigdcd", "ejdhfhhchb", "aecjgajdbffeagahfc", "daaiheffhehcaiha", "gbjfbghecdbbcecj", "ghhahhffj", "fefdjg", "dijgbfjgbhfgafbicdgd", "cjhicaddjffgeeahd", "aaegbcgidchffhgiejgi", "hjcdaahbgfja", "aajidhfhfeg", "hhdageejf", "begiifeabcc", "aiaajehfb", "fegefijiheijhjejic", "d", "ghfadhbaegdeaaihghh", "h", "bjdadfcegjjcgjhjai", "iiifjiejeedibec", "adbcagchbhfcbdid", "hjgj", "afbhbidhh", "bjcchefiahghh", "ggefcifj", "gceddbfafebfegca", "adfgbhaiheehehabfbi", "ffjdfbabdadjf", "gajigbefifdbhfga", "jfhigiaaabffahhagbhc", "iffgiajcdfje", "ehjgafhbhcdjcechf", "bh", "ddajichiacfij", "jcbcjjaafhheahbejadf", "jagggdijggfacabcj", "ihgchcge", "bfbjgjfeg", "djica", "jcjidadjhd", "ahjajdhjhdgeaaehfh", "dhcjifcghjccjd", "bbj", "idbgejhcdg", "begfd", "edeajebcihghef", "ejhjh", "eceajiehdabg", "gdheijaeichacegdfa", "fhfafdddg", "jiibjdidbcfg", "gjjhagjjbgdbjiaf", "bb", "fcehjhfgjbhje", "jcgbcaghcbg", "afhegij", "fhbggicjjdjg", "cehdiibhibcicebfd", "ciabfjfgffcjgbjbjcd", "jdiei", "ebdhaahichcff", "gbejdachjdfec", "faahgjiddibbhdhgi", "aafeeadbdfcdibjdj", "ahaffe", "gjihhgegfh", "cgeededdifbbhcg", "bgf", "cahafcgfidgjgajajcb", "gacdhab", "id", "ccgeb", "hgiifeaiiaafcbh", "ffdcgfjggijhdggddhb", "idhhb", "gaif", "hefjahejhhhhhe", "dj", "hggeegcaiichegfedchf", "bc", "geiadbdgcaeii", "cgbggfbadcbfbhidhjg", "chijihchhgef", "edacjbjfjgbdecddg", "ggjjdddchdhii", "jigbd", "ifbajff", "ihdcfihcjfj", "ddbefaefa", "edbhgbcbhjgaeheeijgh", "ghcdjehfgb", "chfjdgjcdffj", "ffh", "dgc", "iaggidcacghjgahfeg", "dddbhbdcggbb", "baiiicibgej", "iaichdghc", "eddcdjaffebjdafagcie", "eiiachijhcf", "aajifgbdi", "hg", "bibjfg", "ibfdeaafbdfd", "idebdfffigifhiebjb", "jchgeffabfei", "eghjdcei", "jiafgdjadcchfe", "cbjbeejfihgdfheeb", "aiec", "biicjgcd", "dcfeh", "iehehiagjeeh", "bdj", "fi", "fage", "afjechgehgfchb", "ehia", "fadieabh", "bccaiidbiffbebedb", "gajjgaf", "ffii", "iidicfb", "ea", "ejdcfacacjhiffa", "ef", "hhhddgjaaidgdigc", "cicfcieiigiicif", "hah", "ahijijhhcddeabafcgj", "eefcicg", "gbfcifhiegaghhc", "faa", "gbihh", "idfdffif", "ifaijgdcefi", "ddfdbadhgabf", "gghhifiba", "cijadagjhjjacfif", "ibeffhdhhejfgebbdbf", "eeifidadeag", "aadcfjbcbegici", "eaaccfdd", "gdeacedgjbagb", "eahbdccfffgj", "hdceaaiacih", "fagcgfejbcfgacbai", "ghe", "fjdcahhdgcigcchgfhbh", "cafig", "jieebdi", "cjfhdajiejdhhig", "ahejhbjjjdeijcb", "bfbcjjccgi", "bedeji", "gajbgiafdh", "hjifeieiccgjjejei", "jdddjcjcjhjihdha", "hijheac", "dfffigfagibbca", "iacijdfjbafcjhffa", "cdeihee", "dgiaheaeii", "ehahbb", "jeeigigefaj", "fgaghcgbehfgeheacih", "aijbcbgibfbbcfjfffba", "abbaiaagcigjbdific", "bbacahj", "jdiehhddjjeiijgge", "hhigcgibhbdg", "aachgbagja", "edhb", "ajccifc", "fcdcjihjhaacdfehfchj", "bjcg", "bibbdbg", "jdjihegfdde", "hdjf"}
	t.Log(palindromePairs(words))
	t.Log(palindromePairs2(words))
}

// 超时，时间复杂度N^2
// 将每一对i j构成一个新字符串判断其是不是回文
// 时间复杂度 N^2 * M
func palindromePairs(words []string) [][]int {
	m := make(map[string]bool)
	ret := make([][]int, 0)
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			tmp := canPalindrome(words[i], words[j], i, j, m)
			if len(tmp) != 0 {
				ret = append(ret, tmp)
			}
		}
	}
	return ret
}

func canPalindrome(s1, s2 string, k1, k2 int, m map[string]bool) []int {
	if s1+s2 == "" {
		return []int{k1, k2}
	}
	if v, ok := m[s1+s2]; ok {
		if v {
			return []int{k1, k2}
		}
	} else {
		m[s1+s2] = isPalindrome(s1 + s2)
		if m[s1+s2] {
			return []int{k1, k2}
		}
	}
	return nil
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// 对于一个字符串(较长)，它可以和另外一个字符串组成回文时，它可以被分隔成两部分，pre和suf，另一个字符串的翻转时pre或suf，另一个本身为回文
// 时间复杂度M^2 * N
func palindromePairs2(words []string) [][]int {
	ret := make([][]int, 0)
	wordsMap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		wordsMap[words[i]] = i
	}
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i])+1; j++ {
			pre, suf := words[i][:j], words[i][j:]                                  // 出现重复解，寻找长度等于它自身长度的匹配时
			if v, ok := wordsMap[reverse(pre)]; ok && v != i && isPalindrome(suf) { // pre为翻转，suf为回文
				ret = append(ret, []int{i, v})
			}
			// 在suf为翻转时加上j > 0剔除j=0的情况，j=0时pre为空，suf为其本身，避免重复解(在上一个if中加上j < len(words[i]同理)
			if v, ok := wordsMap[reverse(suf)]; ok && v != i && j > 0 && isPalindrome(pre) { // pre为回文，suf为翻转
				ret = append(ret, []int{v, i})
			}
		}
	}
	return ret
}

func reverse(s string) string {
	b := []byte(s)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return string(b)
}
