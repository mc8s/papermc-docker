package last_builds

type LastBuildsParser interface {
	EnsureExists() error
	GetLastBuilds() (LastBuilds, error)
	SaveLastBuilds(lastBuilds LastBuilds) error
}
