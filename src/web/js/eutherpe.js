function closeAddToPlaylist() {
    openConfig("Collection");
}

function closeAddTags() {
    openConfig("Collection");
}

function closeRemoveTags() {
    openConfig("Collection");
}

function closePlayByTags() {
    openConfig("Collection");
}

function showRemoveTagsDiv() {
    getCommonTags();
}

function showPlayByTags() {
    openConfig("PlayByTags");
}

function showAddTags() {
    songSelection = document.getElementsByClassName("CollectionSong");
    selectedOnes = getSelectedSongs(songSelection);
    if (selectedOnes.length == 0) {
        return;
    }
    openConfig("AddTags");
}

function delTagsFromSelection() {
    unselectedOnes = getUnselectedElements("Tag");
    var tagsToDelete = "";
    for (u = 0; u < unselectedOnes.length; u++) {
        tagsToDelete += unselectedOnes[u];
        if ((u + 1) < unselectedOnes.length) {
            tagsToDelete += ",";
        }
    }
    doEutherpeRequest("/eutherpe", { "action" : "collection-untagselections",
                                     "selection" : JSON.stringify(getLastSelection()),
                                     "tags" : tagsToDelete }, "post", true);
    closeRemoveTags();
}

function getLastSelection() {
    return document.getElementById("lastSelection").value.split(",");
}

function addToNext() {
    metaCollectionAdd("collection-addselectiontonext");
}

function addToUpNext() {
    metaCollectionAdd("collection-addselectiontoupnext");
}

function addToPlaylist() {
    metaCollectionAdd("collection-addselectiontoplaylist");
    closeAddToPlaylist();
}

function addTagsToSelection() {
    metaActionOverSongSelection("collection-tagselectionas", "CollectionSong");
}

function playByGivenTags() {
    doEutherpeRequest("/eutherpe", { "action" : "collection-playbygiventags",
                                     "tags" :  getTagContext(),
                                     "amount" : document.getElementById("songsAmount").value }, "post", true);
}

function getTagList() {
    return document.getElementById("tagsSet").value;
}

function getTagContext() {
    return document.getElementById("tagsCtx").value;
}

function getCommonTags() {
    metaActionOverSongSelection("get-commontags", "CollectionSong");
}

function metaCollectionAdd(action) {
    metaActionOverSongSelection(action, "CollectionSong");
    clearCollectionSelection();
}

function clearCollectionSelection() {
    setUnchecked(document.getElementsByClassName("CollectionArtist"));
    setUnchecked(document.getElementsByClassName("CollectionAlbum"));
    setUnchecked(collectionSongs);
}

function clearPlaylistSelection() {
    setUnchecked(document.getElementsByClassName("PlaylistName"));
    setUnchecked(document.getElementsByClassName("PlaylistSong"));
    setUnchecked(collectionSongs);
}

function getSelectedSongs(songList) {
    selectedOnes = [];
    for (var s = 0; s < songList.length; s++) {
        if (songList[s].checked) {
            selectedOnes.push(songList[s].id);
        }
    }
    return selectedOnes;
}

function setUnchecked(list) {
    for (var l = 0; l < list.length; l++) {
        if (list[l].checked) {
            list[l].checked = false;
        }
    }
}

function selectSingleElement(sender) {
    allElements = document.getElementsByClassName(sender.className);
    for (var a = 0; a < allElements.length; a++) {
        allElements[a].checked = false;
    }
    sender.checked = true;
}

function getSelectedBluetoothDevice() {
    return getSelectedElement("BluetoothDevice");
}

function selectPlaylist(sender) {
    var state = sender.checked;
    setUnchecked(document.getElementsByClassName("PlaylistSong"));
    playlists = document.getElementsByClassName("Playlist");
    for (var p = 0; p < playlists.length; p++) {
        playlists[p].checked = false;
    }
    sender.checked = state;
}

function getSelectedPlaylist() {
    return getSelectedElement("Playlist");
}

function getSelectedStorageDevice() {
    return getSelectedElement("StorageDevice");
}

function pairDevice() {
    blueDev = getSelectedBluetoothDevice();
    if (blueDev === null) {
        alert("You must pick up one device.");
        return;
    }
    if (!confirm("Are you sure you want to pair with this device?")) {
        return;
    }
    doEutherpeRequest("/eutherpe", { "action" : "bluetooth-pair",
                                     "bluetooth-device" : blueDev.id }, "post");
}

function unpairDevice() {
    blueDev = getSelectedBluetoothDevice();
    if (blueDev === null) {
        alert("You must pick up one device.");
        return;
    }
    if (!confirm("Are you sure you want to unpair with this device?")) {
        return;
    }
    doEutherpeRequest("/eutherpe", { "action" : "bluetooth-unpair",
                                     "bluetooth-device" : blueDev.id }, "post");
}

function trustDevice() {
    blueDev = getSelectedBluetoothDevice();
    if (blueDev === null) {
        alert("You must pick up one device.");
        return;
    }
    if (!confirm("Are you sure you want to trust this device?")) {
        return;
    }
    doEutherpeRequest("/eutherpe", { "action" : "bluetooth-trust",
                                     "bluetooth-device" : blueDev.id }, "post");
}

function untrustDevice() {
    blueDev = getSelectedBluetoothDevice();
    if (blueDev === null) {
        alert("You must pick up one device.");
        return;
    }
    if (!confirm("Are you sure you want to untrust this device?")) {
        return;
    }
    doEutherpeRequest("/eutherpe", { "action" : "bluetooth-untrust",
                                     "bluetooth-device" : blueDev.id }, "post");
}

function probeDevices() {
    doEutherpeRequest("/eutherpe", { "action" : "bluetooth-probedevices",
                                     "sleep" : "30" }, "post");
}

function listStorageDevices() {
    doEutherpeRequest("/eutherpe", { "action" : "storage-list" }, "post");
}

function scanStorageDevice() {
    doEutherpeRequest("/eutherpe", { "action" : "storage-scan" }, "post");
}

function setStorageDevice() {
    storageDev = getSelectedStorageDevice();
    doEutherpeRequest("/eutherpe", { "action" : "storage-set",
                                     "storage-device" : storageDev.id }, "post");
}

function removePlaylist() {
    playlist = getSelectedPlaylist();
    if (playlist === null) {
        alert("You must pick up one playlist.");
        return;
    }
    if (!confirm("Are you sure you want to remove the playlist '" + playlist.id + "'?")) {
        return;
    }
    doEutherpeRequest("/eutherpe", { "action" : "playlist-remove",
                                     "playlist" : playlist.id }, "post");
}

function clearAllPlaylist() {
    playlist = getSelectedPlaylist();
    if (playlist === null) {
        alert("You must pick up one playlist.");
        return;
    }
    if (!confirm("Are you sure you want to empty the playlist '" + playlist.id + "'?")) {
        return;
    }
    doEutherpeRequest("/eutherpe", { "action" : "playlist-clearall",
                                     "playlist" : playlist.id }, "post");
}

function createPlaylist(playlist) {
    doEutherpeRequest("/eutherpe", { "action" : "playlist-create", "playlist" : playlist }, "post");
}

function moveUpPlaylistSongs() {
    metaActionPlaylistSongs("playlist-moveup");
}

function moveDownPlaylistSongs() {
    metaActionPlaylistSongs("playlist-movedown");
}

function removeSelectedSongsFromPlaylist() {
    if (!confirm("Are you sure?")) {
        return;
    }
    metaActionPlaylistSongs("playlist-removesongs");
    clearPlaylistSelection();
}

function reproducePlaylist() {
    playlist = getSelectedPlaylist();
    if (playlist === null) {
        alert("You must pick one playlist.");
        return;
    }
    doEutherpeRequest("/eutherpe", { "action" : "playlist-reproduce",
                                     "playlist" : playlist.id }, "post");
}

function reproduceSelectedOnesFromPlaylist() {
    metaActionPlaylistSongs("playlist-reproduceselectedones");
    clearPlaylistSelection();
}

function removeSelectedSongsFromUpNext() {
    metaActionMusic("music-remove", "UpNext");
}

function moveSelectedSongsFromUpNextUp() {
    metaActionMusic("music-moveup", "UpNext");
}

function moveSelectedSongsFromUpNextDown() {
    metaActionMusic("music-movedown", "UpNext");
}

function clearUpNextAll() {
    doEutherpeRequest("/eutherpe", { "action" : "music-clearall" }, "post");
}

function shuffleUpNext() {
    doEutherpeRequest("/eutherpe", { "action" : "music-shuffle" }, "post");
}

function musicPlayOrStop(sender) {
    var action = (sender.value != "\u25A0") ? "music-play" : "music-stop";
    doEutherpeRequest("/eutherpe", { "action" : action }, "post");
}

function musicNext() {
    doEutherpeRequest("/eutherpe", { "action" : "music-next" }, "post");
}

function musicLast() {
    doEutherpeRequest("/eutherpe", { "action" : "music-last" }, "post");
}

function musicRepeatAll() {
    doEutherpeRequest("/eutherpe", { "action" : "music-repeatall" }, "post");
}

function musicRepeatOne() {
    doEutherpeRequest("/eutherpe", { "action" : "music-repeatone" }, "post");
}

function getSelectedElement(className) {
    objects = document.getElementsByClassName(className);
    for (var o = 0; o < objects.length; o++) {
        if (objects[o].checked) {
            return objects[o];
        }
    }
    return null;
}

function getUnselectedElements(className) {
    var unselectedOnes = [];
    objects = document.getElementsByClassName(className);
    for (var o = 0; o < objects.length; o++) {
        if (!objects[o].checked) {
            unselectedOnes.push(objects[o].id);
        }
    }
    return unselectedOnes;
}


function metaActionPlaylistSongs(action) {
    metaActionOverSongSelection(action, "PlaylistSong");
}

function metaActionMusic(action) {
    metaActionOverSongSelection(action, "UpNext");
}

function metaActionOverSongSelection(action, songListClassName) {
    songSelection = document.getElementsByClassName(songListClassName);
    selectedOnes = getSelectedSongs(songSelection);
    if (selectedOnes.length == 0) {
        return;
    }
    var reqParams = { "action"    : action,
                      "selection" : JSON.stringify(selectedOnes) };
    if (action == "collection-addselectiontoplaylist") {
        reqParams.playlist = document.getElementById("playlistName").value;
    } else if (songListClassName == "PlaylistSong") {
        playlist = getSelectedPlaylist();
        if (playlist === null) {
            alert("No playlist was selected!");
            return;
        }
        reqParams.playlist = playlist.id;
    } else if (action == "collection-tagselectionas") {
        reqParams.tags = getTagList();
    }
    doEutherpeRequest("/eutherpe", reqParams, "post");
}

function setButtonLabel(glyph, label) {
    if (window.matchMedia("(orientation: portrait)").matches) {
        document.write(glyph);
    } else {
        document.write(glyph + " " + label);
    }
}

function setUncheckedAllSongsOutFromPlaylist(sender) {
    e = sender.id.indexOf(":");
    if (e == -1) {
        return;
    }
    playlistOfThisSong = sender.id.substring(0, e);
    playlists = document.getElementsByClassName("Playlist");
    for (var p = 0; p < playlists.length; p++) {
        playlists[p].checked = (playlists[p].id == playlistOfThisSong);
    }
    playlistSongs = document.getElementsByClassName("PlaylistSong");
    for (var p = 0; p < playlistSongs.length; p++) {
        if (!playlistSongs[p].id.startsWith(playlistOfThisSong)) {
            playlistSongs[p].checked = false;
        }
    }
}

function requestPlayerStatus() {
    if (document.getElementById("Music").style.display != "block") {
        return;
    }
    var req = new XMLHttpRequest();
    req.open("POST", "/eutherpe", true);
    req.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    req.onreadystatechange = function () {
        if (req.readyState == 4 && req.status == 200) {
            try {
                response = JSON.parse(req.responseText);
                nowPlayingDiv = document.getElementById("NowPlayingDiv");
                nowPlayingDiv.innerHTML = response["now-playing"];
                albumCover = document.getElementById("albumCover");
                if (albumCover != null) {
                    albumCover.src = response["album-cover-src"];
                }
                playStop = document.getElementById("playStop");
                if (response["now-playing"].length == 0) {
                    playStop.value = "\u25BA";
                } else if (playStop.value != "\u25A0") {
                    location.replace(location.href);
                }
            } catch (e) {
            }
        }
    };
    req.onerror = function() {
    };
    req.send("action=player-status");
}

function setVolume() {
    doEutherpeRequest("/eutherpe", { "action" : "music-setvolume",
                                     "volume-level" : document.getElementById("volumeLevel").value }, "post", true);
}

function doEutherpeRequest(vdoc, userData, method, noWaitBanner = false) {
    var form = document.createElement("form");
    form.method = method;
    form.action = vdoc;
    for (let k in userData) {
        if (userData.hasOwnProperty(k)) {
            const field = document.createElement("input");
            field.type = "hidden";
            field.name = k;
            field.value = userData[k];
            form.appendChild(field);
        }
    }
    document.body.appendChild(form);
    if (!noWaitBanner) {
        openConfig("Loading");
    }
    form.submit();
}
