package main

var agdaSymbols = [][]string{
	{"∼", "≄", "⋪", "⋡", "⨆", "⌈", "∭", "↹", "⟼", "➾", "↕", "┘", "╜", "╀", "┇", "▣", "◠", "𝟡", "⁆", "¼", "⚄", "⍆", "⍯", "ρ", "𝑌", "𝑵", "𝒞", "𝓐", "𝓸", "𝔥", "③", "➒", "⒝", "Ⓞ"},
	{"∽", "≇", "⋬", "∈", "⋓", "⌉", "∮", "↢", "⟾", "⊸", "⇕", "├", "╟", "┦", "┉", "▢", "◡", "(", "⁾", "½", "⚅", "⍇", "⍰", "Ρ", "𝑍", "𝑶", "𝒟", "𝓑", "𝓹", "𝔦", "⒊", "⑽", "Ⓑ", "ⓞ"},
	{"≈", "≆", "⋠", "∉", "∐", "⌊", "∯", "↩", "⤇", "↑", "↨", "┤", "╢", "┟", "┋", "▤", "◴", "[", "₎", "¾", "⁎", "⍈", "⍱", "σ", "𝑎", "𝑷", "𝒢", "𝓒", "𝓺", "𝔧", "❸", "⑩", "ⓑ", "⒫"},
	{"≋", "≢", ">", "∊", "⨿", "⌋", "∰", "↫", "↷", "⇑", "⇳", "┬", "╥", "╁", "◂", "▥", "◵", "{", ">", "⅓", "⁑", "⍉", "⍲", "Σ", "𝑏", "𝑸", "𝒥", "𝓓", "𝓻", "𝔨", "➂", "⒑", "⒞", "Ⓟ"},
	{"∻", "≭", "≫", "∋", "⊽", "∎", "∱", "⇋", "↻", "⟰", "↔", "┼", "╫", "┧", "◃", "▦", "◶", "⁅", "⎵", "⅔", "⁂", "⍊", "⍳", "τ", "𝑐", "𝑹", "𝒦", "𝓔", "𝓼", "𝔩", "➌", "❿", "Ⓒ", "ⓟ"},
	{"∾", "<", "⋙", "∌", "⊻", "×", "∲", "⇜", "⟳", "⇈", "⇔", "┴", "╨", "┢", "◄", "▧", "◷", "⁽", "⟆", "⅕", "✢", "⍋", "⍴", "Τ", "𝑑", "𝑺", "𝒩", "𝓕", "𝓽", "𝔪", "⑷", "➉", "ⓒ", "⒬"},
	{"∿", "≪", "≥", "∍", "⊍", "∘", "∳", "⇤", "⇰", "⇅", "⇼", "╴", "━", "╈", "◅", "▨", "⚆", "₍", "⟧", "⅖", "✣", "⍌", "⍵", "υ", "𝑒", "𝑻", "𝒪", "𝓖", "𝓾", "𝔫", "④", "➓", "⒟", "Ⓠ"},
	{"≀", "⋘", "≧", "⋲", "⨃", "∙", "∟", "⟻", "⇴", "↥", "↭", "╵", "┃", "┪", "▸", "▩", "⚇", "<", "⟩", "⅗", "✤", "⍍", "⍶", "Υ", "𝑓", "𝑼", "𝒫", "𝓗", "𝓿", "𝔬", "⒋", "⑾", "Ⓓ", "ⓠ"},
	{"≃", "≤", "≳", "⋳", "⊎", "⋆", "∡", "⟽", "⟴", "⇧", "⇿", "╶", "┏", "┡", "▹", "◧", "⚈", "⎴", "⟫", "⅘", "✥", "⍎", "⍷", "φ", "𝑔", "𝑽", "𝒬", "𝓘", "𝔀", "𝔭", "❹", "⑪", "ⓓ", "⒭"},
	{"⋍", "≦", "≷", "⋴", "⨄", "∔", "∢", "⤆", "⟿", "↟", "⟷", "╷", "┓", "╇", "►", "◨", "⚉", "⟅", "⦄", "⅙", "✱", "⍏", "⍸", "Φ", "𝑖", "𝑾", "𝒮", "𝓙", "𝔁", "𝔮", "➃", "⒒", "⒠", "Ⓡ"},
	{"≂", "≲", "≻", "⋵", "⊌", "∸", "⊾", "↶", "➵", "↿", "⟺", "╭", "┗", "┩", "▻", "◩", "✦", "⟦", ">", "⅚", "✲", "⍐", "⍹", "χ", "𝑗", "𝑿", "𝒯", "𝓚", "𝔂", "𝔯", "➍", "⑿", "Ⓔ", "ⓡ"},
	{"≅", "≶", "≽", "⋶", "∑", "∶", "⊿", "↺", "➸", "↾", "↮", "╮", "┛", "┮", "▴", "◪", "✧", "⅛", "✳", "⍑", "⍺", "Χ", "𝑘", "𝒀", "𝒰", "𝓛", "𝔃", "𝔰", "⑸", "⑫", "ⓔ", "⒮"},
	{"≌", "≺", "≿", "⋷", "⅀", "∷", "⋈", "⟲", "➙", "⇡", "⇎", "╯", "┣", "┭", "▵", "◫", "✶", "⅜", "✺", "⍒", "⎕", "ψ", "𝑙", "𝒁", "𝒱", "𝓜", "𝔄", "𝔱", "⑤", "⒓", "⒡", "Ⓢ"},
	{"≊", "≼", "⊃", "⋸", "⊢", "∺", "⋉", "→", "➔", "⇞", "⇹", "╰", "┫", "┶", "▾", "◰", "✴", "⅝", "✻", "⍓", "α", "Ψ", "𝑚", "𝒂", "𝒲", "𝓝", "𝔅", "𝔲", "⒌", "⒀", "Ⓕ", "ⓢ"},
	{"≡", "≾", "⊇", "⋹", "⊣", "∹", "⋊", "⇒", "➛", "↰", "↖", "╱", "┳", "┵", "▿", "◱", "✹", "⅞", "✼", "⍔", "Α", "ω", "𝑛", "𝒃", "𝒳", "𝓞", "𝔇", "𝔳", "❺", "⑬", "ⓕ", "⒯"},
	{"≣", "⊂", "⋑", "⋺", "⊤", "⊹", "⋋", "⇛", "➜", "↱", "⇖", "╲", "╋", "┾", "◢", "◲", "★", "⅟", "✽", "⍕", "β", "Ω", "𝑜", "𝒄", "𝒴", "𝓟", "𝔈", "𝔴", "➄", "⒔", "⒢", "Ⓣ"},
	{"≐", "⊆", "⊐", "⋻", "⊥", "∛", "⋌", "⇉", "➝", "➦", "⇱", "╳", "┻", "┽", "◿", "◳", "☆", "•", "❃", "⍖", "Β", "𝐴", "𝑝", "𝒅", "𝒵", "𝓠", "𝔉", "𝔵", "➎", "⒁", "Ⓖ", "ⓣ"},
	{"≑", "⋐", "⊒", "⋼", "⊦", "∜", "⨝", "⇄", "➞", "⇪", "↸", "═", "╸", "┲", "◣", "▬", "✪", "◦", "❉", "⍗", "γ", "𝐵", "𝑞", "𝒆", "𝒶", "𝓡", "𝔊", "𝔶", "⑹", "⑭", "ⓖ", "⒰"},
	{"≒", "⊏", "⊱", "⋽", "⊧", "∆", "⟕", "↦", "➟", "⇫", "↗", "║", "╹", "┱", "◺", "▭", "✫", "‣", "❊", "⍘", "Γ", "𝐶", "𝑟", "𝒇", "𝒷", "𝓢", "𝔍", "𝔷", "⑥", "⒕", "⒣", "Ⓤ"},
	{"≓", "⊑", "⊳", "⋾", "⊨", "∞", "⟖", "⇨", "➠", "⇬", "⇗", "╔", "╺", "┺", "◤", "▮", "✯", "⁌", "❋", "⍙", "δ", "𝐷", "𝑠", "𝒈", "𝒸", "𝓣", "𝔎", "¡", "⒍", "⒂", "Ⓗ", "ⓤ"},
	{"≔", "⊰", "⊵", "⋿", "⊩", "⅋", "⟗", "↠", "➡", "⇭", "➶", "╗", "╻", "┹", "◸", "▯", "✰", "⁍", "✂", "⍚", "Δ", "𝐸", "𝑡", "𝒉", "𝒹", "𝓤", "𝔏", "¢", "❻", "⑮", "ⓗ", "⒱"},
	{"≕", "⊲", "⋗", "∩", "⊪", "⊕", "←", "⇀", "➢", "⇮", "➹", "╚", "┍", "╊", "◥", "▰", "✵", "♩", "✄", "⍛", "ε", "𝐹", "𝑢", "𝒊", "𝒻", "𝓥", "𝔐", "¦", "➅", "⒖", "⒤", "Ⓥ"},
	{"≖", "⊴", "⋛", "⋂", "⊫", "⊖", "⇐", "⇁", "➣", "⇯", "➚", "╝", "┯", "╉", "◹", "▱", "✷", "♪", "⁀", "⍜", "Ε", "𝐺", "𝑣", "𝒋", "𝒽", "𝓦", "𝔑", "°", "➏", "⒃", "Ⓘ", "ⓥ"},
	{"≗", "⋖", "⋝", "∧", "⊬", "⊗", "⇚", "⇢", "➤", "↓", "↘", "╠", "┑", "╆", "◀", "◆", "✸", "♫", "‿", "⍝", "ζ", "𝐻", "𝑤", "𝒌", "𝒾", "𝓧", "𝔒", "¿", "⑺", "⑯", "ⓘ", "⒲"},
	{"≘", "⋚", "⋟", "⋀", "⊭", "⊘", "⇇", "⇻", "➧", "⇓", "⇘", "╣", "┕", "╅", "◁", "◇", "ℕ", "♬", "⌶", "⍞", "Ζ", "𝐼", "𝑥", "𝒍", "𝒿", "𝓨", "𝔓", "ª", "⑦", "⒗", "⒥", "Ⓦ"},
	{"≙", "⋜", "≯", "⋏", "⊮", "⊙", "⇆", "↝", "➨", "⟱", "⇲", "╦", "┷", "╄", "▶", "◈", "ℤ", "♭", "⌷", "⍠", "θ", "𝐽", "𝑦", "𝒎", "𝓀", "𝓩", "𝔔", "º", "⒎", "⒄", "Ⓙ", "ⓦ"},
	{"≚", "⋞", "≱", "⨇", "⊯", "⊚", "↤", "⇾", "➩", "⇊", "➴", "╬", "┙", "╃", "▷", "●", "ℚ", "♯", "⌸", "⍡", "Θ", "𝐾", "𝑧", "𝒏", "𝓁", "𝓪", "𝔖", "⓪", "❼", "⑰", "ⓙ", "⒳"},
	{"≛", "≮", "≩", "⊓", "∣", "⊛", "⇦", "⟶", "➪", "⇵", "➷", "╩", "┝", "╿", "▲", "○", "ℝ", "-", "-", "ι", "𝐿", "𝑨", "𝒐", "𝓂", "𝓫", "𝔗"},
	{"≜", "≰", "≵", "⨅", "∤", "⊜", "↞", "⟹", "➫", "↧", "➘", "╒", "┿", "╽", "△", "◎", "ℂ", "-", ")", "–", "⌺", "⍣", "Ι", "𝑀", "𝑩", "𝒑", "𝓃", "𝓬", "𝔘", "①", "➐", "⒅", "Ⓚ", "ⓧ"},
	{"≝", "≨", "⋧", "⋒", "∥", "⊝", "↼", "↛", "➬", "⇩", "↙", "╕", "┥", "╼", "▼", "◌", "ℙ", "-", ")", "—", "⌻", "⍤", "κ", "𝑁", "𝑪", "𝒒", "𝓅", "𝓭", "𝔙", "⒈", "⑻", "⑱", "ⓚ", "⒴"},
	{"≞", "≴", "≹", "∏", "∦", "⨁", "↽", "⇏", "➭", "↡", "⇙", "╘", "┎", "╾", "▽", "◯", "𝔹", "(", "]", "ⁱ", "⌼", "⍥", "Κ", "𝑂", "𝑫", "𝒓", "𝓆", "𝓮", "𝔚", "❶", "⑧", "⒙", "⒧", "Ⓨ"},
	{"≟", "⋦", "⊁", "⊼", "∀", "⨂", "⇠", "⇸", "➮", "⇃", "⋯", "╛", "┰", "╌", "◬", "◍", "𝟘", "(", "}", "‼", "⌽", "⍦", "λ", "𝑃", "𝑬", "𝒔", "𝓇", "𝓯", "𝔛", "➀", "⒏", "⒆", "Ⓛ", "ⓨ"},
	{"≍", "≸", "⋩", "⨉", "∃", "⨀", "⇺", "⇶", "➯", "⇂", "⋮", "╞", "┒", "╎", "◭", "◐", "𝟙", "(", "｣", "⁇", "⌾", "⍧", "Λ", "𝑄", "𝑭", "𝒕", "𝓈", "𝓰", "𝔜", "➊", "❽", "⑲", "ⓛ", "⒵"},
	{"≎", "⊀", "⊅", "∪", "∄", "⍟", "↜", "↴", "➱", "⇣", "⋰", "╡", "┖", "┄", "◮", "◑", "𝟚", "(", "′", "‽", "⌿", "⍨", "ƛ", "𝑅", "𝑮", "𝒖", "𝓉", "𝓱", "𝔞", "⑵", "➇", "⒚", "⒨", "Ⓩ"},
	{"≏", "⋨", "⊉", "⋃", "∅", "⊞", "⇽", "↣", "➲", "⇟", "⋱", "╤", "┸", "┆", "■", "◒", "𝟛", "[", "″", "⁈", "⍀", "⍩", "μ", "𝑆", "𝑯", "𝒗", "𝓊", "𝓲", "𝔟", "②", "➑", "⒇", "Ⓜ", "ⓩ"},
	{"≬", "⊄", "⊋", "∨", "∁", "⊟", "⟵", "↪", "➳", "↵", "─", "╪", "┚", "┈", "□", "◓", "𝟜", "{", "‴", "⁉", "⍁", "⍪", "Μ", "𝑇", "𝑰", "𝒘", "𝓋", "𝓳", "𝔠", "⒉", "⑼", "⑳", "ⓜ"},
	{"⋕", "⊈", "⋣", "⋁", "⌜", "⊠", "⟸", "↬", "➺", "↲", "│", "╧", "┠", "┊", "◼", "◔", "𝟝", "｢", "⁗", "⚀", "⍂", "⍫", "ν", "𝑈", "𝑱", "𝒙", "𝓌", "𝓴", "𝔡", "❷", "⑨", "⒛", "⒩"},
	{"≠", "⊊", "⋥", "⋎", "⌝", "⊡", "↚", "⇌", "➻", "↳", "┌", "╓", "╂", "╍", "◻", "◕", "𝟞", ")", "‵", "⚁", "⍃", "⍬", "Ν", "𝑉", "𝑲", "𝒚", "𝓍", "𝓵", "𝔢", "➁", "⒐", "⒜", "Ⓝ"},
	{"≁", "⋢", "⋫", "⨈", "⌞", "∫", "⇍", "⇝", "➼", "➥", "┐", "╖", "┨", "╏", "-", "◖", "𝟟", "]", "‶", "⚂", "⍄", "⍭", "ξ", "𝑊", "𝑳", "𝒛", "𝓎", "𝓶", "𝔣", "➋", "❾", "Ⓐ", "ⓝ"},
	{"≉", "⋤", "⋭", "⊔", "⌟", "∬", "⇷", "⇥", "➽", "↯", "└", "╙", "┞", "┅", "-", "◗", "𝟠", "}", "‷", "⚃", "⍅", "⍮", "Ξ", "𝑋", "𝑴", "𝒜", "𝓏", "𝓷", "𝔤", "⑶", "➈", "ⓐ", "⒪"},
}
