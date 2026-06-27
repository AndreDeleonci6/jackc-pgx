func (p *Pool) Acquire(ctx context.Context) (*Conn, error) {
	for {
		resource, err := p.puddle.Acquire(ctx)
		if err != nil {
			return nil, err
		}

		conn := resource.Value().(*Conn)

		if p.config.BeforeAcquire != nil {
			ok, err := p.config.BeforeAcquire(ctx, conn)
			if err != nil || !ok {
				resource.Destroy()
				continue
			}
		}

		return conn, nil
	}
}