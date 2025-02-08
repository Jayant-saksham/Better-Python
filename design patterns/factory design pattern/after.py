"""
Video exporting app
"""

from abc import ABC, abstractmethod

class VideoExporter(ABC):

    @abstractmethod
    def prepare_export(self, video_data):
        pass

    @abstractmethod
    def do_export(self, folder: str):
        pass

class LosslessVideoExporter(VideoExporter):
    def prepare_export(self, video_data):
        print("Preparing video data for lossless export")

    def do_export(self, folder: str):
        print(f"Exporting video data in lossless format {folder}")

class MediumVideoExporter(VideoExporter):

    def prepare_export(self, video_data):
        print("Preparing video data for medium export")

    def do_export(self, folder: str):
        print(f"Exporting video data in medium format {folder}")

class LowVideoExporter(VideoExporter):

    def prepare_export(self, video_data):
        print("Preparing video data for low export")

    def do_export(self, folder: str):
        print(f"Exporting video data in low format {folder}")


class AudioExporter(ABC):

    @abstractmethod
    def prepare_export(self, video_data):
        pass

    @abstractmethod
    def do_export(self, folder: str):
        pass

class HighAudioExporter(AudioExporter):
    def prepare_export(self, video_data):
        print("Preparing audio data for high export")

    def do_export(self, folder: str):
        print(f"Exporting audio data in high format to {folder}")

class MediumAudioExporter(AudioExporter):
    def prepare_export(self, video_data):
        print("Preparing audio data for medium export")

    def do_export(self, folder: str):
        print(f"Exporting audio data in medium format to {folder}")

class LowAudioExporter(AudioExporter):
    def prepare_export(self, video_data):
        print("Preparing audio data for low export")

    def do_export(self, folder: str):
        print(f"Exporting audio data in low format to {folder}")

class ExporterFactory(ABC):
    @abstractmethod
    def get_video_exporter(self):
        pass

    @abstractmethod
    def get_audio_exporter(self):
        pass

class HighQualityExporterFactory(ExporterFactory):
    def get_audio_exporter(self):
        return HighAudioExporter()

    def get_video_exporter(self):
        return LosslessVideoExporter()


class MediumQualityExporterFactory(ExporterFactory):
    def get_audio_exporter(self):
        return MediumAudioExporter()

    def get_video_exporter(self):
        return MediumVideoExporter()


class LowQualityExporterFactory(ExporterFactory):
    def get_audio_exporter(self):
        return LowAudioExporter()

    def get_video_exporter(self):
        return LowVideoExporter()

def get_object():
    factories = {
        "high": HighQualityExporterFactory(),
        "low": LowQualityExporterFactory(),
        "medium": MediumQualityExporterFactory()
    }

    user_input = input("Enter input (high, low, medium) : ")
    if user_input in factories:
        return factories[user_input]


if __name__ == "__main__":
    factory = get_object()
    audio_object = factory.get_audio_exporter()
    video_object = factory.get_video_exporter()

    audio_object.do_export("Folder")
    video_object.do_export("Folder")