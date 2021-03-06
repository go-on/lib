package html

import "github.com/go-on/lib/types"

// entity list stolen from http://unicode.e-workers.de/entities.php

var (
	E_Acirc    = types.HTMLString(`&Acirc;`)
	E_acirc    = types.HTMLString(`&acirc;`)
	E_acute    = types.HTMLString(`&acute;`)
	E_AElig    = types.HTMLString(`&AElig;`)
	E_aelig    = types.HTMLString(`&aelig;`)
	E_Agrave   = types.HTMLString(`&Agrave;`)
	E_agrave   = types.HTMLString(`&agrave;`)
	E_alefsym  = types.HTMLString(`&alefsym;`)
	E_Alpha    = types.HTMLString(`&Alpha;`)
	E_alpha    = types.HTMLString(`&alpha;`)
	E_amp      = types.HTMLString(`&amp;`)
	E_and      = types.HTMLString(`&and;`)
	E_ang      = types.HTMLString(`&ang;`)
	E_apos     = types.HTMLString(`&apos;`)
	E_Aring    = types.HTMLString(`&Aring;`)
	E_aring    = types.HTMLString(`&aring;`)
	E_asymp    = types.HTMLString(`&asymp;`)
	E_Atilde   = types.HTMLString(`&Atilde;`)
	E_atilde   = types.HTMLString(`&atilde;`)
	E_Auml     = types.HTMLString(`&Auml;`)
	E_auml     = types.HTMLString(`&auml;`)
	E_bdquo    = types.HTMLString(`&bdquo;`)
	E_Beta     = types.HTMLString(`&Beta;`)
	E_beta     = types.HTMLString(`&beta;`)
	E_brvbar   = types.HTMLString(`&brvbar;`)
	E_bull     = types.HTMLString(`&bull;`)
	E_cap      = types.HTMLString(`&cap;`)
	E_Ccedil   = types.HTMLString(`&Ccedil;`)
	E_ccedil   = types.HTMLString(`&ccedil;`)
	E_cedil    = types.HTMLString(`&cedil;`)
	E_cent     = types.HTMLString(`&cent;`)
	E_Chi      = types.HTMLString(`&Chi;`)
	E_chi      = types.HTMLString(`&chi;`)
	E_circ     = types.HTMLString(`&circ;`)
	E_clubs    = types.HTMLString(`&clubs;`)
	E_cong     = types.HTMLString(`&cong;`)
	E_copy     = types.HTMLString(`&copy;`)
	E_crarr    = types.HTMLString(`&crarr;`)
	E_cup      = types.HTMLString(`&cup;`)
	E_curren   = types.HTMLString(`&curren;`)
	E_Dagger   = types.HTMLString(`&Dagger;`)
	E_dagger   = types.HTMLString(`&dagger;`)
	E_dArr     = types.HTMLString(`&dArr;`)
	E_darr     = types.HTMLString(`&darr;`)
	E_deg      = types.HTMLString(`&deg;`)
	E_Delta    = types.HTMLString(`&Delta;`)
	E_delta    = types.HTMLString(`&delta;`)
	E_diams    = types.HTMLString(`&diams;`)
	E_divide   = types.HTMLString(`&divide;`)
	E_Eacute   = types.HTMLString(`&Eacute;`)
	E_eacute   = types.HTMLString(`&eacute;`)
	E_Ecirc    = types.HTMLString(`&Ecirc;`)
	E_ecirc    = types.HTMLString(`&ecirc;`)
	E_Egrave   = types.HTMLString(`&Egrave;`)
	E_egrave   = types.HTMLString(`&egrave;`)
	E_empty    = types.HTMLString(`&empty;`)
	E_emsp     = types.HTMLString(`&emsp;`)
	E_ensp     = types.HTMLString(`&ensp;`)
	E_Epsilon  = types.HTMLString(`&Epsilon;`)
	E_epsilon  = types.HTMLString(`&epsilon;`)
	E_equiv    = types.HTMLString(`&equiv;`)
	E_Eta      = types.HTMLString(`&Eta;`)
	E_eta      = types.HTMLString(`&eta;`)
	E_ETH      = types.HTMLString(`&ETH;`)
	E_eth      = types.HTMLString(`&eth;`)
	E_Euml     = types.HTMLString(`&Euml;`)
	E_euml     = types.HTMLString(`&euml;`)
	E_euro     = types.HTMLString(`&euro;`)
	E_exist    = types.HTMLString(`&exist;`)
	E_fnof     = types.HTMLString(`&fnof;`)
	E_forall   = types.HTMLString(`&forall;`)
	E_frac12   = types.HTMLString(`&frac12;`)
	E_frac14   = types.HTMLString(`&frac14;`)
	E_frac34   = types.HTMLString(`&frac34;`)
	E_frasl    = types.HTMLString(`&frasl;`)
	E_Gamma    = types.HTMLString(`&Gamma;`)
	E_gamma    = types.HTMLString(`&gamma;`)
	E_ge       = types.HTMLString(`&ge;`)
	E_gt       = types.HTMLString(`&gt;`)
	E_hArr     = types.HTMLString(`&hArr;`)
	E_harr     = types.HTMLString(`&harr;`)
	E_hearts   = types.HTMLString(`&hearts;`)
	E_hellip   = types.HTMLString(`&hellip;`)
	E_Iacute   = types.HTMLString(`&Iacute;`)
	E_iacute   = types.HTMLString(`&iacute;`)
	E_Icirc    = types.HTMLString(`&Icirc;`)
	E_icirc    = types.HTMLString(`&icirc;`)
	E_iexcl    = types.HTMLString(`&iexcl;`)
	E_Igrave   = types.HTMLString(`&Igrave;`)
	E_igrave   = types.HTMLString(`&igrave;`)
	E_image    = types.HTMLString(`&image;`)
	E_infin    = types.HTMLString(`&infin;`)
	E_int      = types.HTMLString(`&int;`)
	E_Iota     = types.HTMLString(`&Iota;`)
	E_iota     = types.HTMLString(`&iota;`)
	E_iquest   = types.HTMLString(`&iquest;`)
	E_isin     = types.HTMLString(`&isin;`)
	E_Iuml     = types.HTMLString(`&Iuml;`)
	E_iuml     = types.HTMLString(`&iuml;`)
	E_Kappa    = types.HTMLString(`&Kappa;`)
	E_kappa    = types.HTMLString(`&kappa;`)
	E_Lambda   = types.HTMLString(`&Lambda;`)
	E_lambda   = types.HTMLString(`&lambda;`)
	E_lang     = types.HTMLString(`&lang;`)
	E_laquo    = types.HTMLString(`&laquo;`)
	E_lArr     = types.HTMLString(`&lArr;`)
	E_larr     = types.HTMLString(`&larr;`)
	E_lceil    = types.HTMLString(`&lceil;`)
	E_ldquo    = types.HTMLString(`&ldquo;`)
	E_le       = types.HTMLString(`&le;`)
	E_lfloor   = types.HTMLString(`&lfloor;`)
	E_lowast   = types.HTMLString(`&lowast;`)
	E_loz      = types.HTMLString(`&loz;`)
	E_lrm      = types.HTMLString(`&lrm;`)
	E_lsaquo   = types.HTMLString(`&lsaquo;`)
	E_lsquo    = types.HTMLString(`&lsquo;`)
	E_lt       = types.HTMLString(`&lt;`)
	E_macr     = types.HTMLString(`&macr;`)
	E_mdash    = types.HTMLString(`&mdash;`)
	E_micro    = types.HTMLString(`&micro;`)
	E_middot   = types.HTMLString(`&middot;`)
	E_minus    = types.HTMLString(`&minus;`)
	E_Mu       = types.HTMLString(`&Mu;`)
	E_mu       = types.HTMLString(`&mu;`)
	E_nabla    = types.HTMLString(`&nabla;`)
	E_nbsp     = types.HTMLString(`&nbsp;`)
	E_ndash    = types.HTMLString(`&ndash;`)
	E_ne       = types.HTMLString(`&ne;`)
	E_ni       = types.HTMLString(`&ni;`)
	E_not      = types.HTMLString(`&not;`)
	E_notin    = types.HTMLString(`&notin;`)
	E_nsub     = types.HTMLString(`&nsub;`)
	E_Ntilde   = types.HTMLString(`&Ntilde;`)
	E_ntilde   = types.HTMLString(`&ntilde;`)
	E_Nu       = types.HTMLString(`&Nu;`)
	E_nu       = types.HTMLString(`&nu;`)
	E_Oacute   = types.HTMLString(`&Oacute;`)
	E_oacute   = types.HTMLString(`&oacute;`)
	E_Ocirc    = types.HTMLString(`&Ocirc;`)
	E_ocirc    = types.HTMLString(`&ocirc;`)
	E_OElig    = types.HTMLString(`&OElig;`)
	E_oelig    = types.HTMLString(`&oelig;`)
	E_Ograve   = types.HTMLString(`&Ograve;`)
	E_ograve   = types.HTMLString(`&ograve;`)
	E_oline    = types.HTMLString(`&oline;`)
	E_Omega    = types.HTMLString(`&Omega;`)
	E_omega    = types.HTMLString(`&omega;`)
	E_Omicron  = types.HTMLString(`&Omicron;`)
	E_omicron  = types.HTMLString(`&omicron;`)
	E_oplus    = types.HTMLString(`&oplus;`)
	E_or       = types.HTMLString(`&or;`)
	E_ordf     = types.HTMLString(`&ordf;`)
	E_ordm     = types.HTMLString(`&ordm;`)
	E_Oslash   = types.HTMLString(`&Oslash;`)
	E_oslash   = types.HTMLString(`&oslash;`)
	E_Otilde   = types.HTMLString(`&Otilde;`)
	E_otilde   = types.HTMLString(`&otilde;`)
	E_otimes   = types.HTMLString(`&otimes;`)
	E_Ouml     = types.HTMLString(`&Ouml;`)
	E_ouml     = types.HTMLString(`&ouml;`)
	E_para     = types.HTMLString(`&para;`)
	E_part     = types.HTMLString(`&part;`)
	E_permil   = types.HTMLString(`&permil;`)
	E_perp     = types.HTMLString(`&perp;`)
	E_Phi      = types.HTMLString(`&Phi;`)
	E_phi      = types.HTMLString(`&phi;`)
	E_Pi       = types.HTMLString(`&Pi;`)
	E_pi       = types.HTMLString(`&pi;`)
	E_piv      = types.HTMLString(`&piv;`)
	E_plusmn   = types.HTMLString(`&plusmn;`)
	E_pound    = types.HTMLString(`&pound;`)
	E_Prime    = types.HTMLString(`&Prime;`)
	E_prime    = types.HTMLString(`&prime;`)
	E_prod     = types.HTMLString(`&prod;`)
	E_prop     = types.HTMLString(`&prop;`)
	E_Psi      = types.HTMLString(`&Psi;`)
	E_psi      = types.HTMLString(`&psi;`)
	E_quot     = types.HTMLString(`&quot;`)
	E_radic    = types.HTMLString(`&radic;`)
	E_rang     = types.HTMLString(`&rang;`)
	E_raquo    = types.HTMLString(`&raquo;`)
	E_rArr     = types.HTMLString(`&rArr;`)
	E_rarr     = types.HTMLString(`&rarr;`)
	E_rceil    = types.HTMLString(`&rceil;`)
	E_rdquo    = types.HTMLString(`&rdquo;`)
	E_real     = types.HTMLString(`&real;`)
	E_reg      = types.HTMLString(`&reg;`)
	E_rfloor   = types.HTMLString(`&rfloor;`)
	E_Rho      = types.HTMLString(`&Rho;`)
	E_rho      = types.HTMLString(`&rho;`)
	E_rlm      = types.HTMLString(`&rlm;`)
	E_rsaquo   = types.HTMLString(`&rsaquo;`)
	E_rsquo    = types.HTMLString(`&rsquo;`)
	E_sbquo    = types.HTMLString(`&sbquo;`)
	E_Scaron   = types.HTMLString(`&Scaron;`)
	E_scaron   = types.HTMLString(`&scaron;`)
	E_sdot     = types.HTMLString(`&sdot;`)
	E_sect     = types.HTMLString(`&sect;`)
	E_shy      = types.HTMLString(`&shy;`)
	E_Sigma    = types.HTMLString(`&Sigma;`)
	E_sigma    = types.HTMLString(`&sigma;`)
	E_sigmaf   = types.HTMLString(`&sigmaf;`)
	E_sim      = types.HTMLString(`&sim;`)
	E_spades   = types.HTMLString(`&spades;`)
	E_sub      = types.HTMLString(`&sub;`)
	E_sube     = types.HTMLString(`&sube;`)
	E_sum      = types.HTMLString(`&sum;`)
	E_sup      = types.HTMLString(`&sup;`)
	E_sup1     = types.HTMLString(`&sup1;`)
	E_sup2     = types.HTMLString(`&sup2;`)
	E_sup3     = types.HTMLString(`&sup3;`)
	E_supe     = types.HTMLString(`&supe;`)
	E_szlig    = types.HTMLString(`&szlig;`)
	E_Tau      = types.HTMLString(`&Tau;`)
	E_tau      = types.HTMLString(`&tau;`)
	E_there4   = types.HTMLString(`&there4;`)
	E_Theta    = types.HTMLString(`&Theta;`)
	E_theta    = types.HTMLString(`&theta;`)
	E_thetasym = types.HTMLString(`&thetasym;`)
	E_thinsp   = types.HTMLString(`&thinsp;`)
	E_THORN    = types.HTMLString(`&THORN;`)
	E_thorn    = types.HTMLString(`&thorn;`)
	E_tilde    = types.HTMLString(`&tilde;`)
	E_times    = types.HTMLString(`&times;`)
	E_trade    = types.HTMLString(`&trade;`)
	E_Uacute   = types.HTMLString(`&Uacute;`)
	E_uacute   = types.HTMLString(`&uacute;`)
	E_uArr     = types.HTMLString(`&uArr;`)
	E_uarr     = types.HTMLString(`&uarr;`)
	E_Ucirc    = types.HTMLString(`&Ucirc;`)
	E_ucirc    = types.HTMLString(`&ucirc;`)
	E_Ugrave   = types.HTMLString(`&Ugrave;`)
	E_ugrave   = types.HTMLString(`&ugrave;`)
	E_uml      = types.HTMLString(`&uml;`)
	E_upsih    = types.HTMLString(`&upsih;`)
	E_Upsilon  = types.HTMLString(`&Upsilon;`)
	E_upsilon  = types.HTMLString(`&upsilon;`)
	E_Uuml     = types.HTMLString(`&Uuml;`)
	E_uuml     = types.HTMLString(`&uuml;`)
	E_weierp   = types.HTMLString(`&weierp;`)
	E_Xi       = types.HTMLString(`&Xi;`)
	E_xi       = types.HTMLString(`&xi;`)
	E_Yacute   = types.HTMLString(`&Yacute;`)
	E_yacute   = types.HTMLString(`&yacute;`)
	E_yen      = types.HTMLString(`&yen;`)
	E_Yuml     = types.HTMLString(`&Yuml;`)
	E_yuml     = types.HTMLString(`&yuml;`)
	E_Zeta     = types.HTMLString(`&Zeta;`)
	E_zeta     = types.HTMLString(`&zeta;`)
	E_zwj      = types.HTMLString(`&zwj;`)
	E_zwnj     = types.HTMLString(`&zwnj;`)
)
