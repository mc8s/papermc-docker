package last_builds

import "time"

type LastBuilds struct {
	LastBuilds []LastBuild `json:"last_builds"`
}

type LastBuild struct {
	Project string `json:"project"`
	Tag     string `json:"tag"`
	Build   string `json:"build"`
	Date    string `json:"date"`
}

func (l *LastBuilds) GetLastBuild(project string, tag string) string {
	for _, lastBuild := range l.LastBuilds {
		if lastBuild.Project == project && lastBuild.Tag == tag {
			return lastBuild.Build
		}
	}
	return ""
}

func (l *LastBuilds) AddLastBuild(project string, tag string, build string) {
	newBuild := LastBuild{
		Project: project,
		Tag:     tag,
		Build:   build,
		Date:    time.Now().String(),
	}
	if l.GetLastBuild(project, tag) == "" {
		l.LastBuilds = append(l.LastBuilds, newBuild)
	}
	for i, lastBuild := range l.LastBuilds {
		if lastBuild.Project == project && lastBuild.Tag == tag {
			l.LastBuilds[i] = newBuild
		}
	}
}
