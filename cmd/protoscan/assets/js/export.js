const Export = (function () {

    function Export() {
    }

    const b64toBlob = function (b64Data, contentType, sliceSize) {
        // function taken from http://stackoverflow.com/a/16245768/2591950
        // author Jeremy Banks http://stackoverflow.com/users/1114/jeremy-banks
        contentType = contentType || '';
        sliceSize = sliceSize || 512;

        const byteCharacters = window.atob(b64Data);
        const byteArrays = [];

        let offset;
        for (offset = 0; offset < byteCharacters.length; offset += sliceSize) {
            const slice = byteCharacters.slice(offset, offset + sliceSize);

            const byteNumbers = new Array(slice.length);
            let i;
            for (i = 0; i < slice.length; i = i + 1) {
                byteNumbers[i] = slice.charCodeAt(i);
            }

            const byteArray = new window.Uint8Array(byteNumbers);

            byteArrays.push(byteArray);
        }

        return new window.Blob(byteArrays, {
            type: contentType
        });
    };

    const createDownloadLink = function (anchor, base64data, exporttype, filename) {
        if (window.navigator.msSaveBlob) {
            const blob = b64toBlob(base64data, exporttype);
            window.navigator.msSaveBlob(blob, filename);
            return false;
        } else if (window.URL.createObjectURL) {
            const blob = b64toBlob(base64data, exporttype);
            anchor.href = window.URL.createObjectURL(blob);
        } else {
            anchor.download = filename;
            anchor.href = "data:" + exporttype + ";base64," + base64data;
        }

        // Return true to allow the link to work
        return true;
    };

    Export.createDownloadLink = function (anchor, filename, base64data) {
        return createDownloadLink(anchor, base64data, 'application/octet-stream', filename);
    };

    return Export;
})();