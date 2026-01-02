
package system

func GetSystemStats() (SystemStats, error) {
	var s SystemStats
	var err error

	s.CPUUsage, err = CPUUsage()
	if err != nil {
		return s, err
	}

	s.MemTotalMB, s.MemUsedMB, s.MemUsagePct, err = MemoryUsage()
	if err != nil {
		return s, err
	}

	s.Load1, s.Load5, s.Load15, err = LoadAverage()
	if err != nil {
		return s, err
	}

	s.DiskUsage, err = DiskUsage("/")
	if err != nil {
		return s, err
	}

	s.Uptime, err = Uptime()
	if err != nil {
		return s, err
	}

	return s, nil
}
