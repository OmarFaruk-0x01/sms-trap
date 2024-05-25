function isASCII(str) {
    if (!str) {
        return false;
    }

    for (let i = 0; i < str.length; i++) {
        if (str.trim().charCodeAt(i) > 127) {
            return false;
        }
    }

    return true;
}

function isUnicode(str) {
    if (!str) {
        return false;
    }

    for (let i = 0; i < str.length; i++) {
        if (str.trim().charCodeAt(i) > 127) {
            return true;
        }
    }

    return false;
}

function detectEncoding(str) {
    if (isASCII(str)) {
        return "Text";
    }

    if (isUnicode(str)) {
        return "Unicode";
    }

    return "";
}

function countGSMChars(str) {
    if (!str) {
        return 0;
    }

    const gsmChars = '@£$¥èéùìòÇ\nØø\rÅåΔ_ΦΓΛΩΠΨΣΘΞ\x1BÆæßÉ !"#¤%&\'()*+,-./0123456789:;<=>?¡ABCDEFGHIJKLMNOPQRSTUVWXYZÄÖÑÜ§¿abcdefghijklmnopqrstuvwxyzäöñüà';

    let gsmCount = 0;
    for (let i = 0; i < str.length; i++) {
        if (gsmChars.includes(str[i])) {
            gsmCount++;
        }
    }
    return gsmCount;
}

function countUnicodeChars(str) {
    if (!str) {
        return 0;
    }

    let unicodeCount = 0;
    for (let i = 0; i < str.length; i++) {
        if (isUnicode(str[i])) {
            unicodeCount++;
        }
    }
    return unicodeCount;
}