package main

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func srvrestarter(serviceName string) error {
	os := runtime.GOOS

	switch os {
	case "windows":
		err := restartExternalServiceWindows(serviceName)
		if err != nil {
			return err
		}
	case "linux":
		err := restartExternalServiceLinux(serviceName)
		if err != nil {
			return err
		}
	case "darwin":
		err := restartExternalServiceMacOS(serviceName)
		if err != nil {
			return err
		}
	case "freebsd":
		err := restartExternalServiceFreeBSD(serviceName)
		if err != nil {
			return err
		}
	default:
		err := errors.New("Перезапуск службы в ОС не предусмотрен\n")
		return err
	}
	return nil
}

func restartExternalServiceWindows(serviceName string) error {

	// Остановить службу
	cmdStop := exec.Command("net", "stop", serviceName)
	if err := cmdStop.Run(); err != nil {
		logger.Printf("Не удалось остановить службу %s: %v\n", serviceName, err)
		// Возможно, служба уже остановлена, попробуем запустить
	}

	// Дать время на остановку
	time.Sleep(2 * time.Second)

	// Запустить службу
	cmdStart := exec.Command("net", "start", serviceName)
	if err := cmdStart.Run(); err != nil {
		return fmt.Errorf("не удалось запустить службу %s: %w", serviceName, err)
	}
	logger.Printf("Служба %s успешно перезапущена\n", serviceName)

	return nil
}

func restartExternalServiceLinux(serviceName string) error {

	// Выполняем команду: sudo systemctl restart <serviceName>
	cmd := exec.Command("sudo", "systemctl", "restart", serviceName)

	err := cmd.Run()
	if err != nil {
		logger.Fatalf("Ошибка при перезапуске службы: %v", err)
		return fmt.Errorf("не удалось перезапустить службу %s: %w", serviceName, err)
	}

	logger.Printf("Служба %s успешно перезапущена", serviceName)

	return nil
}

func restartExternalServiceMacOS(serviceName string) error {

	// Остановить службу
	cmdStop := exec.Command("sudo", "launchctl", "stop", serviceName)
	if err := cmdStop.Run(); err != nil {
		fmt.Printf("Не удалось остановить службу %s: %v\n", serviceName, err)
		// Возможно, служба уже остановлена, попробуем запустить
	}

	// Дать время на остановку
	time.Sleep(2 * time.Second)

	// Запустить службу
	cmdStart := exec.Command("sudo", "launchctl", "start", serviceName)
	if err := cmdStart.Run(); err != nil {
		return fmt.Errorf("не удалось запустить службу %s: %w", serviceName, err)
	}
	fmt.Printf("Служба %s успешно перезапущена\n", serviceName)

	return nil
}

func restartExternalServiceFreeBSD(serviceName string) error {

	cmd := exec.Command("sudo", "service", serviceName, "restart")

	err := cmd.Run()
	if err != nil {
		logger.Fatalf("Ошибка при перезапуске службы: %v", err)
		return fmt.Errorf("не удалось перезапустить службу %s: %w", serviceName, err)
	}

	logger.Printf("Служба %s успешно перезапущена", serviceName)

	return nil
}
