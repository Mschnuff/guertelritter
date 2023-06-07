# Kanal- und Programmdaten aus PyWebOSTV

Im folgenden wird der Rückgabe-Inhalt aus den von der Bibliothek [PyWebOSTV](https://github.com/supersaiyanmode/PyWebOSTV) bereitgestellten Funktionen erläutert. Die Tabellen sollen nach und nach vervollständigt bzw. korrigiert werden.
  



## get_current_channel_program_info
------
Funktion stammt direkt aus der API und war nicht in der Bibliothek enthalten. Infos über das aktuelle Programm und den aktuellen Kanal.

   | Schlüssel | Wert/Bedeutung |
   |---------| ---------------- |   
   | "programId" | __Identifikationsnummer__ (Zusammensetzung unklar) |
   | "programName" | __Name der Sendung__ |
   | "description" | __Zusatzinfos (z.B Herkunftsland, Genre etc.)__ |
   | "startTime" | __Startzeit nach GMT (Format: Jahr,Monat,Tag,Stunde,Minuten,Sekunden)__ |
   | "endTime" | __Endzeit nach GMT (Format: Jahr,Monat,Tag,Stunde,Minuten,Sekunden)__ |
   | "localStartTime" | __wie 'startTime' nur GMT+1__ |
   | "localEndTime" | __wie 'endTime' nur GMT+1__ |
   | "duration" | __Dauer der Sendung in Sekunden__ |
   | "channelId" | __Identifikationsnummer des Kanals__ (1_25_channelNumber_[signalChannelId](#get_channel_data)) |
   | "channelName" | __Name des Kanals__ |
   | "channelNumber" | __Identifikationsnummer des Kanals (wie auf dem Fernseher)__ |
   | "channelMode" | __terrestrisch oder digital__ |
  



## get_channel_data
------
Infos über den aktuellen Kanal

   | Schlüssel | Wert/Bedeutung |
   |---------| ---------------- |
   | "channelId" | __Identifikationsnummer des Kanals__ (1_25_channelNumber_[signalChannelId](#get_channel_data)) |
   | "physicalNumber" | ? (bei ARD: 25) |
   | "isScrambled" | __verschlüsselt__ (CI-Modul wird benötigt) |
   | "channelTypeName" | __DVB-T, DVB-S, DVB-C__ oder __DVB-H__ |
   | "isLocked" | __Kanäle können mit 4stelliger Pin gesperrt werden (Kindersicherung)__ |
   | "dualChannel" | zwei Fernsehsignale gleichzeitig empfangen? |
   | "isChannelChanged" | ? |
   | "channelModeName" | __terrestrisch oder digital__ (als Wort) |
   | "channelNumber" | __Identifikationsnummer des Kanals (wie auf dem Fernseher)__ |
   | "isFineTuned" | __moderne Fernseher können bei analogen Kanälen die Frequenz nachtunen__ |
   | "channelTypeId" | ? (bei ARD: 1) |
   | "isDescrambled" | __entschlüsselt / unverschlüsselt__ (nicht gleich mit "descrambled" von [get_program_data](#get_program_data) |
   | "isSkipped" | __Kanäle in der Skipped-Liste werden beim Sender Suchlauf ignoriert__ |
   | "isHEVCChannel" | __H.265 /MPEG-H__ (Videokodierungsstandard) |
   | "hybridtvType" | __HBBTV oder was anderes (außerhalb der EU)__ |
   | "isInvisible" | ? |
   | "favoriteGroup" | __ist Kanal in der Favoriten-Liste enhalten__ |
   | "channelName" | __Name des Kanals__ (z.B. "Das Erste HD") |
   | "channelModeId" | __terrestrisch oder digital__ (als Zahl) |
   | "signalChannelId" | __[TSID](#get_program_data)_[SVCID](#get_program_data)_8468__ |




## get_program_data 
------
Programmvorschau-Liste und weitere Infos über den aktuellen Kanal

#### channel
Hierin sind nocheinmal ein Paar Infos zum aktuellen Kanal enthalten. Einige davon redundant.
   | Schlüssel | Wert/Bedeutung |
   |---------| ---------------- |
   | "channelId" | __Identifikationsnummer des Kanals__ (1_25_channelNumber_[signalChannelId](#get_channel_data)) |
   | "programId" | __Identifikationsnummer des Programms__ (wie [channelId](#get_channel_data)) |
   | "signalChannelId" | __[TSID](#get_program_data)_[SVCID](#get_program_data)_8468__ |
   | "chanCode" | ? (bei ARD: "UNKNOWN") |
   | "channelMode" | __terrestrisch oder digital__ |
   | "channelModeId" | __terrestrisch oder digital__ (als Zahl) |
   | "channelType" | __DVB-T, DVB-S, DVB-C__ oder __DVB-H__ |
   | "channelTypeId": | __DVB-T, DVB-S, DVB-C__ oder __DVB-H__ als Zahl |
   | "channelNumber": |  [__Identifikationsnummer des Kanals (wie auf dem Fernseher)__](#get_channel_data) |
   | "majorNumber": | ? (bei ARD: 1) |
   | "minorNumber": | ? (bei ARD: 0) |
   | "channelName" | __Name des Kanals__ |
   | "skipped" | __Kanäle in der Skipped-Liste werden beim Sender Suchlauf ignoriert__ |
   | "locked" |  __Kanäle können mit 4stelliger Pin gesperrt werden__ (Kindersicherung) |
   | "descrambled" |  __entschlüsselt / unverschlüsselt__ (nicht gleich mit "isDescrambled" von [get_channel_data](#get_channel_data) |
   | "scrambled" | __verschlüsselt__ (CI-Modul wird benötigt) |
   | "serviceType" | ? (bei ARD: 16) |
   | "favoriteGroup" | ? (bei ARD: leer) | 
   | "imgUrl" | URL des Kanal-Logos? |
   | "display" | Kanal mit Bild? (bei ARD: 1) |
   | "satelliteName" | Feld ist leer | 
   | "fineTuned" | __moderne Fernseher können bei analogen Kanälen die Frequenz nachtunen__ |
   | Frequency" | __Senderfrequenz__ |
   | "shortCut" | Existiert ein shortcut zum Sender? (bei ARD: 0) | 
   | "Bandwidth" | ? (bei ARD: 0) |
   | "HDTV" | __High Definition Television__ (Vertikalauflösung 720p oder 1080p) |
   | "Invisible" | ? |
   | "TV" | __ist ein TV-Sender__ |
   | "DTV" | __Kanal benötigt DVBT-Decoder__ |
   | "ATV" | [?](https://www.techtarget.com/whatis/definition/ATV-advanced-television) |
   | "Data" | ? |
   | "Radio" | __ist ein Radio-Sender__ |
   | "Numeric" | ? |
   | "PrimaryCh" | Primärkanal? |
   | "specialService" | ? |
   | "CASystemIDList" | ? (bei ARD: leer) |       
   | "CASystemIDListCount | ? (bei ARD: 0) |
   | "groupIdList" | ? (ei ARD: [ 0, 1 ]) |
   | "channelGenreCode" | Genre des Kanals? (bei ARD: "UNKNOWN") |
   | "favoriteIdx[A-bis-H]" | __[IDX](https://en.wikipedia.org/wiki/IDX)__ (bei ARD: 250) |
   | "imgUrl2" | URL des Kanal-Logos? |
   | "channelLogoSize" | Bildgröße von [waterMarkUrl](#get_program_data)? (bei ARD: "UNKNOWN") |
   | "ipChanServerUrl" | [__Server-URL bei Digital TV__](https://support.encodedmedia.com/article/101-adding-external-channels-to-your-tv-server) |
   | "payChan" | __kostenpflichtiger Kanal__ |
   | "IPChannelCode" | ? (bei ARD: "UNKNOWN") |
   | "ipCallNumber" | ? (bei ARD: "UNKNOWN") |
   | "outFlag" | ? (bei ARD: false) |
   | "satelliteLcn" | [__ist 'Logical Channel Numbering' anhand vom Satelliten übertragener Kanalnummer aktiviert__](https://wikideck.com/de/Logical_Channel_Numbering) |
   | "waterMarkUrl" | __vermutlich die Url des Programmlogos__ |
   | "channelNameSortKey" | ? (bei ARD: "20") |
   | "ipChanType" | ? (bei ARD: "UNKNOWN") |
   | "adultFlag" | Kanal für Erwachsene ( . )( . ) |
   | "ipChanCategory" | [?](https://en.wikipedia.org/wiki/Internet_Protocol_television) (bei ARD: "UNKNOWN") |
   | "ipChanInteractive" | [?](https://en.wikipedia.org/wiki/Internet_Protocol_television) (bei ARD: false) |
   | "callSign" | ? (bei ARD: "UNKNOWN") |
   | "adFlag" | __Kanal mit Werbung?__ |
   | "configured" | ? (bei ARD: false) |
   | "lastUpdated" | bei gespeicherten Kanälen: Datum der letzten Änderung? |
   | "ipChanCpId"| control protocol? (bei ARD: "UNKNOWN") |
   | "isFreeviewPlay" | ? (bei ARD: 0) |
   | "hasBackward" | ? (bei ARD: 0) |
   | "playerService" | [__Software des vom Fernseher genutzten Players?__](https://www.webosose.org/docs/reference/) (vermutlich immer: "com.webos.service.tv") |
   | "TSID" | [__TSID__](https://de.wikipedia.org/wiki/Transpondertabelle) |
   | "SVCID" | __Signalling Virtual Channel Identifier Downstream__ (keine Ahnung was das ist... ) |  
  



#### programList
Die Program-Liste enthält eine Vorschau mit Sendungen des nächsten Tages.
Anzahl der Sendungen hängt vermutlich ab von Dauer der einzelnen Sendungen.

Je Sendung:  
| Schlüssel | Wert/Bedeutung |
|---------| ---------------- |  
| "channelId" | __Identifikationsnummer des Kanals__ (1_25_channelNumber_[signalChannelId](#get_channel_data)) |
| "duration" | __Dauer der Sendung in Sekunden__ |
| "endTime" | __Endzeit nach GMT (Format: Jahr,Monat,Tag,Stunde,Minuten,Sekunden)__ |
| "localStartTime" | __Startzeitpunkt der Sendung nach GMT+1__ |
| "localEndTime" | __wie 'endTime' nur GMT+1__ |
| "genre" | __Genre der Sendung__ |       
| "programId" | __8468_[TSID](#get_program_data)_[SVCID](#get_program_data)_XXXXX__
| "programName" | __Name der Sendung__
| "rating" | __Bewertung der Sendung__ (enthält "ratingString", "ratingValue", "region" und "_id") |
| "signalChannelId" | __[TSID](#get_program_data)_[SVCID](#get_program_data)_8468__ |
| "startTime" | __Startzeit nach GMT (Format: Jahr,Monat,Tag,Stunde,Minuten,Sekunden)__ |
| "tableId" | evtl. kann man sich mehrere Programmlisten zurückgeben lassen ? |
  


## get_channel_id
------
Von mir erstellte Funktion, die direkt die Kanalnummer zurückgibt

return: __Kanalnummer als Zahl__
  

    

## get_info
------
Infos über Hard- und Software des Fernsehgeräts

| Schlüssel | Wert/Bedeutung |
|---------| ---------------- | 
| "product_name": __Name des Betriebssystems__ |
| "model_name": Hardware ID? (bei uns: "HE_DTV_W18A_AFADABAA") |
| "sw_type" | __Typ der Software__ (bei uns: "FIRMWARE") |
| "major_ver" | __Versionsnummer der Firmware (übergeordnete Nummer)__ |  
| "minor_ver" | __Versionsnummer der Firmware (untergeordnete Nummer)__ |  
| "country" | __Staatenkürzel__ (bei uns: "DE") |
| "country_group" | __Regionskürzel__ (bei uns: "EU") |
| "device_id" | __vermutlich MAC-Adresse__ (bei uns: "a8:23:fe:ed:dc:ea") |
| "auth_flag" | ? (bei uns: "N") |
| "ignore_disable" | ? (bei uns: "N") |
| "eco_info" | ? (bei uns: "01") |
| "config_key" | ? (bei uns: "01") |
| "language_code" | __Sprach Code__ (bei uns: "de-DE") |