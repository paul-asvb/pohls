package resolver

import "regexp"

var (
	seasonPattern        = regexp.MustCompile(`(?i)s?([0-9]{1,2})[ex]`)
	episodePattern       = regexp.MustCompile(`(?i)[ex]([0-9]{1,2})`)
	seasonEpisodePattern = regexp.MustCompile(`(?i)s?([0-9]{1,2})[ex]([0-9]{2})(?:[^0-9])`)
	yearPattern          = regexp.MustCompile(`([\[\(]?((?:19[0-9]|20[01])[0-9])[\]\)]?)`)
	resolutionPattern    = regexp.MustCompile(`([0-9]{3,4}p)`)
	qualityPattern       = regexp.MustCompile(`(?i)((?:PPV\.)?[HP]DTV|(?:HD)?CAM|B[DR]Rip|(?:HD-?)?TS|(?:PPV )?WEB-?DL(?: DVDRip)?|HDRip|DVDRip|DVDRIP|CamRip|W[EB]B[Rr]ip|BluRay|DvDScr|hdtv|telesync|HEVC|SNHD)`)
	codecPattern         = regexp.MustCompile(`(?i)(xvid|[hx]\.?26[45])`)
	audioPattern         = regexp.MustCompile(`(?i)(MP3|DD5\.?1|Dual[\- ]Audio|LiNE|DTS|AAC[.-]LC|AAC(?:\.?2\.0)?|AC3(?:\.5\.1)?|(5\.1))`)
	groupPattern         = regexp.MustCompile(`(- ?([^-]+(?:-={[^-]+-?$)?))$`)
	regionPattern        = regexp.MustCompile(`R[0-9]`)
	extendedPattern      = regexp.MustCompile(`(EXTENDED(:?.CUT)?)`)
	hardcodedPattern     = regexp.MustCompile(`HC`)
	properPattern        = regexp.MustCompile(`PROPER`)
	repackPattern        = regexp.MustCompile(`REPACK`)
	containerPattern     = regexp.MustCompile(`(MKV|AVI|MP4)`)
	widescreenPattern    = regexp.MustCompile(`WS`)
	websitePattern       = regexp.MustCompile(`^(\[ ?([^\]]+?) ?\])`)
	languagePattern      = regexp.MustCompile(`(?i)(rus\.eng|ita\.eng|hindi)`)
	sbsPattern           = regexp.MustCompile(`(?:Half-)?SBS`)
	unratedPattern       = regexp.MustCompile(`UNRATED`)
	sizePattern          = regexp.MustCompile(`(?i)(\d+(?:\.\d+)?(?:GB|MB))`)
	is3dPattern          = regexp.MustCompile(`3D`)
	otherPatterns        = regexp.MustCompile(`(?i)READNFO|TiTAN|MAXPRO|UpScaled|iNTERNAL|DesiSCR|Mafiaking|Rip`)
)
